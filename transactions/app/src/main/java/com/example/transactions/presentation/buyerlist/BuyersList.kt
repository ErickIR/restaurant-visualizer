package com.example.transactions.presentation.buyerlist

import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.core.content.ContextCompat
import androidx.core.os.bundleOf
import androidx.navigation.findNavController
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.RecyclerView
import androidx.swiperefreshlayout.widget.SwipeRefreshLayout
import com.example.transactions.R
import com.example.transactions.data.ApiResponse
import com.example.transactions.data.ApiServiceClient
import com.example.transactions.data.PaginatedApiResponse
import com.example.transactions.models.Buyer
import com.google.android.material.datepicker.CalendarConstraints
import com.google.android.material.datepicker.DateValidatorPointBackward
import com.google.android.material.datepicker.MaterialDatePicker
import com.google.android.material.floatingactionbutton.FloatingActionButton
import com.google.android.material.progressindicator.CircularProgressIndicator
import com.google.android.material.snackbar.Snackbar
import com.google.android.material.textfield.TextInputEditText
import com.google.android.material.textfield.TextInputLayout
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response
import java.util.*
import kotlin.collections.ArrayList

/**
 * A simple [Fragment] subclass as the default destination in the navigation.
 */
class BuyersList : Fragment() {
    private val repo = ApiServiceClient()
    private var page = 1
    private var nextPage = 2
    private lateinit var buyersList: ArrayList<Buyer>
    private lateinit var adapter: BuyerListAdapter
    private lateinit var recyclerView: RecyclerView
    private lateinit var linearLayoutManager: LinearLayoutManager
    private lateinit var loadingIndicator: CircularProgressIndicator
    private lateinit var searchTextView: TextInputLayout
    private lateinit var searchEditText: TextInputEditText
    private lateinit var swipeRefreshLayout: SwipeRefreshLayout

    override fun onCreateView(
            inflater: LayoutInflater, container: ViewGroup?,
            savedInstanceState: Bundle?
    ): View? {
        // Inflate the layout for this fragment
        val view = inflater.inflate(R.layout.buyer_list, container, false)
        recyclerView = view.findViewById(R.id.buyersRecyclerView)
        linearLayoutManager = LinearLayoutManager(view.context)
        recyclerView.layoutManager = linearLayoutManager
        buyersList = ArrayList()
        adapter = BuyerListAdapter(buyersList)
        recyclerView.adapter = adapter
        loadingIndicator = view.findViewById(R.id.loadingIndicator)
        searchTextView = view.findViewById(R.id.searchTextInputLayout)
        swipeRefreshLayout = view.findViewById(R.id.swipeRefreshLayout)
        searchEditText = view.findViewById(R.id.searchEditText)
        swipeRefreshLayout.setOnRefreshListener {
            resetBuyers()
            getBuyers(page)
            swipeRefreshLayout.isRefreshing = false
        }

        searchTextView.setEndIconOnClickListener {
            val buyerId: String = searchEditText.text.toString()
            if(buyerId.isEmpty()) {
                showSnackNotification("Type an ID to search...", R.color.red)

            } else {
                val bundle = bundleOf("buyerId" to buyerId)
                view?.findNavController()?.navigate(R.id.actionGoToDetailsPage, bundle)
            }
        }

        recyclerView.addOnScrollListener(object: RecyclerView.OnScrollListener() {
            private var loading = true
            private val visibleThreshold = 5
            private var previousTotal = 0
            override fun onScrolled(recyclerView: RecyclerView, dx: Int, dy: Int) {
                super.onScrolled(recyclerView, dx, dy)

                val visibleItemCount = recyclerView.childCount
                val totalItemCount = linearLayoutManager.itemCount
                val firstVisibleItem = linearLayoutManager.findFirstVisibleItemPosition()

                if (loading) {
                    if (totalItemCount > previousTotal) {
                        loading = false
                        previousTotal = totalItemCount
                    }
                }

                if(!loading && (totalItemCount - visibleItemCount) <= (firstVisibleItem + visibleThreshold)) {
                    println("GETTING MORE DATA")
                    showLoadingIndicator()
                    getBuyers(nextPage)
                    hideLoadingIndicator()
                    loading = true
                }
            }
        })

        setLoadDataDatePicker(view)
        getBuyers(page)

        return view
    }

    private fun resetBuyers() {
        page = 1
        nextPage = 2
        buyersList.clear()
        adapter.notifyDataSetChanged()
    }

    private fun getBuyers(page: Int) {
        val call = repo.getAllBuyers(page)
        call.enqueue(object: Callback<PaginatedApiResponse> {
            override fun onResponse(
                call: Call<PaginatedApiResponse>,
                response: Response<PaginatedApiResponse>
            ) {
                if (response.isSuccessful) {
                    val isSuccess = response.body()?.success ?: false

                    if (isSuccess) {
                        val buyersRequested = response.body()?.data
                        if (buyersRequested?.isNotEmpty() == true) {
                            loadBuyersToList(buyersRequested)
                            this@BuyersList.page = response.body()!!.metadata.page
                            val totalPages = response.body()!!.metadata.totalPages
                            if(page < totalPages){
                                nextPage = page + 1
                            }
                        }
                        if (buyersRequested?.isEmpty() == true) {
                            showSnackNotification("It seems that there is no data on the server.", R.color.red)
                        }
                    } else {
                        val message = response.body()?.message
                        showSnackNotification(message, R.color.red)
                    }
                    hideLoadingIndicator()
                }
            }

            override fun onFailure(call: Call<PaginatedApiResponse>, t: Throwable) {
                val message = t.message
                showSnackNotification(message, R.color.red)
            }
        })
    }

    private fun showSnackNotification(message: String?, resourceColor: Int) {
        Snackbar.make(requireView(), message.toString(), Snackbar.LENGTH_LONG)
            .setTextColor(ContextCompat.getColor(requireContext(), R.color.white))
            .setBackgroundTint(ContextCompat.getColor(requireContext(), resourceColor))
            .show()
    }

    private fun loadBuyersToList(buyers: List<Buyer>){
        buyersList.addAll(buyers)
        adapter.notifyDataSetChanged()
    }

    private fun setLoadDataDatePicker(view: View) {
        val today = MaterialDatePicker.todayInUtcMilliseconds()

        val datePickerBuilder = MaterialDatePicker.Builder.datePicker()
            .setTitleText("SELECT DATE: ")
            .setSelection(today)

        val calendarConstraints =
            CalendarConstraints.Builder()
                .setValidator(DateValidatorPointBackward.now())
                .build()

        val datePicker = datePickerBuilder
            .setCalendarConstraints(calendarConstraints)
            .build()

        datePicker.addOnPositiveButtonClickListener {
            val unixTimeStamp = it / 1000
            showLoadingIndicator()
            loadDataToBd(unixTimeStamp.toString())
        }

        view.findViewById<FloatingActionButton>(R.id.fab).setOnClickListener {
            datePicker.show(parentFragmentManager, "MyTAG")
        }
    }

    private fun showLoadingIndicator() {
        loadingIndicator.visibility = View.VISIBLE
        recyclerView.visibility = View.GONE
    }

    private fun hideLoadingIndicator() {
        recyclerView.visibility = View.VISIBLE
        loadingIndicator.visibility = View.GONE
    }

    private fun loadDataToBd(date: String)  {
        val call = repo.loadData(date)

        call.enqueue(object: Callback<ApiResponse<Unit>> {
            override fun onResponse(
                call: Call<ApiResponse<Unit>>,
                response: Response<ApiResponse<Unit>>
            ) {
                if (response.isSuccessful) {
                    val isSuccess = response.body()?.success ?: false

                    val snackColor = if (isSuccess) R.color.green else R.color.red
                    val message = response.body()?.message

                    showSnackNotification(message, snackColor)
                    resetBuyers()
                    getBuyers(page)
                    hideLoadingIndicator()
                }
            }

            override fun onFailure(call: Call<ApiResponse<Unit>>, t: Throwable) {
                val message = t.message
                showSnackNotification(message, R.color.red)
            }
        })
    }
}