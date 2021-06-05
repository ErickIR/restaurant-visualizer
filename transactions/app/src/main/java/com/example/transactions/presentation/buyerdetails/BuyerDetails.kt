package com.example.transactions.presentation.buyerdetails

import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ImageView
import android.widget.TextView
import androidx.core.content.ContextCompat
import androidx.core.widget.NestedScrollView
import androidx.navigation.fragment.findNavController
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.RecyclerView
import com.example.transactions.R
import com.example.transactions.data.ApiResponse
import com.example.transactions.data.ApiServiceClient
import com.example.transactions.data.dto.BuyerDetailsDto
import com.example.transactions.models.Buyer
import com.example.transactions.models.Product
import com.example.transactions.models.Transaction
import com.example.transactions.presentation.buyerdetails.adapters.OtherBuyersAdapter
import com.example.transactions.presentation.buyerdetails.adapters.ProductstAdapter
import com.example.transactions.presentation.buyerdetails.adapters.TransactionsAdapter
import com.example.transactions.presentation.helper.parseDateFromUnixTimestampToDate
import com.google.android.material.progressindicator.CircularProgressIndicator
import com.google.android.material.snackbar.Snackbar
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response
/**
 * A simple [Fragment] subclass as the second destination in the navigation.
 */
class BuyerDetails : Fragment() {
    private val repo = ApiServiceClient()
    private lateinit var nameTextView: TextView
    private lateinit var idDetailsTextView: TextView
    private lateinit var ageDetailsTextView: TextView
    private lateinit var dateTextView: TextView
    private lateinit var backIconButton: ImageView
    private lateinit var loadingIndicator: CircularProgressIndicator
    private lateinit var nestedScrollView: NestedScrollView
    private lateinit var transactionsTitleLabel: TextView
    private lateinit var otherBuyersTitleLabel: TextView
    private lateinit var recommendedProductsTitleLabel: TextView

    private lateinit var otherBuyersRecyclerView: RecyclerView
    private lateinit var transactionsRecyclerView: RecyclerView
    private lateinit var productRecyclerView: RecyclerView

    private lateinit var productstAdapter: ProductstAdapter
    private lateinit var otherBuyersAdapter: OtherBuyersAdapter
    private lateinit var transactionsAdapter: TransactionsAdapter

    private val otherBuyers = ArrayList<Buyer>()
    private val transactions = ArrayList<Transaction>()
    private val products = ArrayList<Product>()
    private var transactionsExpanded = false
    private var recommendationsExpanded = false
    private var relatedBuyersExpanded = false
    
    private var buyerInfo: BuyerDetailsDto? = null
    override fun onCreateView(
            inflater: LayoutInflater, container: ViewGroup?,
            savedInstanceState: Bundle?
    ): View? {
        // Inflate the layout for this fragment
        val view = inflater.inflate(R.layout.buyer_details, container, false)

        backIconButton = view.findViewById(R.id.backIconButton)
        loadingIndicator = view.findViewById(R.id.loadingIndicator)
        nameTextView = view.findViewById(R.id.nameTextView)
        idDetailsTextView = view.findViewById(R.id.idDetailsTextView)
        ageDetailsTextView = view.findViewById(R.id.ageDetailsTextView)
        dateTextView = view.findViewById(R.id.dateDetailsTextView)
        nestedScrollView = view.findViewById(R.id.nestedScrollView)

        otherBuyersRecyclerView = view.findViewById(R.id.otherBuyersRecyclerView)
        otherBuyersTitleLabel = view.findViewById(R.id.otherBuyersTitleLabel)
        otherBuyersAdapter = OtherBuyersAdapter(otherBuyers)

        transactionsRecyclerView = view.findViewById(R.id.transactionsRecyclerView)
        transactionsTitleLabel = view.findViewById(R.id.transactionsTitleLabel)
        transactionsAdapter = TransactionsAdapter(transactions)

        productRecyclerView = view.findViewById(R.id.recommendedProductsRecyclerView)
        recommendedProductsTitleLabel = view.findViewById(R.id.recommendedProductsTitleLabel)
        productstAdapter = ProductstAdapter(products)

        transactionsRecyclerView.layoutManager = LinearLayoutManager(context)
        transactionsRecyclerView.adapter = transactionsAdapter

        otherBuyersRecyclerView.layoutManager = LinearLayoutManager(context)
        otherBuyersRecyclerView.adapter = otherBuyersAdapter

        productRecyclerView.layoutManager = LinearLayoutManager(context)
        productRecyclerView.adapter = productstAdapter

        otherBuyersTitleLabel.setOnClickListener {
            relatedBuyersExpanded = !relatedBuyersExpanded
            otherBuyersRecyclerView.visibility = if(relatedBuyersExpanded) View.VISIBLE else View.GONE
        }

        transactionsTitleLabel.setOnClickListener {
            transactionsExpanded = !transactionsExpanded
            transactionsRecyclerView.visibility = if(transactionsExpanded) View.VISIBLE else View.GONE
        }

        recommendedProductsTitleLabel.setOnClickListener {
            recommendationsExpanded = !recommendationsExpanded
            productRecyclerView.visibility = if(recommendationsExpanded) View.VISIBLE else View.GONE
        }

        val buyerId = arguments?.getString("buyerId")

        backIconButton.setOnClickListener {
            findNavController().popBackStack()
        }

        getBuyerInformation(buyerId!!)
        return view
    }

    private fun loadInformationIntoView(buyerDetailsDto: BuyerDetailsDto) {
        nameTextView.text = buyerDetailsDto.buyer.name
        idDetailsTextView.text = buyerDetailsDto.buyer.id
        ageDetailsTextView.text = buyerDetailsDto.buyer.age.toString()
        dateTextView.text = buyerDetailsDto.buyer.date?.parseDateFromUnixTimestampToDate() ?: "N/A"

        for (item in buyerDetailsDto.buyersWithSameIp) {
            otherBuyers.add(item.buyer)
        }
        otherBuyersAdapter.notifyDataSetChanged()

        transactions.addAll(buyerDetailsDto.transactions)
        transactionsAdapter.notifyDataSetChanged()

        products.addAll(buyerDetailsDto.products)
        productstAdapter.notifyDataSetChanged()
    }

    private fun getBuyerInformation(buyerId: String)  {
        val call = repo.getBuyerInformation(buyerId)
        call.enqueue(object: Callback<ApiResponse<BuyerDetailsDto>> {
            override fun onResponse(
                call: Call<ApiResponse<BuyerDetailsDto>>,
                response: Response<ApiResponse<BuyerDetailsDto>>
            ) {
                if (response.isSuccessful) {
                    val isSuccess = response.body()?.success ?: false

                    if (isSuccess) {
                        buyerInfo = response.body()?.data
                        nestedScrollView.visibility = View.VISIBLE
                        loadInformationIntoView(buyerInfo!!)
                    } else {
                        val snackColor = R.color.red
                        val message = response.body()?.message

                        showSnackNotification(message, snackColor)
                    }
                    loadingIndicator.visibility = View.GONE


                }
            }

            override fun onFailure(call: Call<ApiResponse<BuyerDetailsDto>>, t: Throwable) {
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
}