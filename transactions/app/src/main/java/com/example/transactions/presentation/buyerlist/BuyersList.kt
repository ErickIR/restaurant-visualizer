package com.example.transactions.presentation.buyerlist

import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.core.content.ContextCompat
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.RecyclerView
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

    override fun onCreateView(
            inflater: LayoutInflater, container: ViewGroup?,
            savedInstanceState: Bundle?
    ): View? {
        // Inflate the layout for this fragment
        val view = inflater.inflate(R.layout.fragment_first, container, false)
        recyclerView = view.findViewById(R.id.buyersRecyclerView)
        buyersList = ArrayList()
        linearLayoutManager = LinearLayoutManager(view.context)
        recyclerView.layoutManager = linearLayoutManager
        adapter = BuyerListAdapter(buyersList)
        recyclerView.adapter = adapter
        loadingIndicator = view.findViewById(R.id.loadingIndicator)

        recyclerView.addOnScrollListener(object: RecyclerView.OnScrollListener() {
            override fun onScrolled(recyclerView: RecyclerView, dx: Int, dy: Int) {
                super.onScrolled(recyclerView, dx, dy)
                if (dy > 0) {
                    val visibleItemCount = linearLayoutManager.childCount
                    val totalItemCount = linearLayoutManager.itemCount
                    val pastVisibleItems = linearLayoutManager.findFirstVisibleItemPosition()

                    if ((visibleItemCount + pastVisibleItems) >= totalItemCount) {
                        showLoadingIndicator()
                        getBuyers(nextPage)
                        hideLoadingIndicator()
                    }
                }
            }
        })

        setLoadDataDatePicker(view)
        getBuyers(page)

        return view
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

                    if (isSuccess) {
                        val message = response.body()?.message
                        showSnackNotification(message, R.color.green)
                    } else {
                        val message = response.body()?.message
                        showSnackNotification(message, R.color.red)
                    }
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