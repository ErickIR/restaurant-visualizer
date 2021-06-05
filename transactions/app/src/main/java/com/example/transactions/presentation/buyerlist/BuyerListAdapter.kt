package com.example.transactions.presentation.buyerlist

import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.core.os.bundleOf
import androidx.navigation.findNavController
import androidx.recyclerview.widget.RecyclerView
import com.example.transactions.R
import com.example.transactions.models.Buyer
import com.example.transactions.presentation.helper.parseDateFromUnixTimestampToDate
import java.util.*
import kotlin.collections.ArrayList

class BuyerListAdapter(private val buyers: ArrayList<Buyer>)
    : RecyclerView.Adapter<BuyerListAdapter.BuyerListViewHolder>() {

    class BuyerListViewHolder(itemView: View) : RecyclerView.ViewHolder(itemView), View.OnClickListener {
        private var buyer: Buyer? = null
        private lateinit var nameTextView: TextView
        private lateinit var ageTextView: TextView
        private lateinit var idTextView: TextView
        private lateinit var dateTextView: TextView
        init {
            itemView.setOnClickListener(this)
        }

        override fun onClick(view: View?) {
            val bundle = bundleOf("buyerId" to buyer?.id)
            view?.findNavController()?.navigate(R.id.actionGoToDetailsPage, bundle)
        }

        fun bindBuyerInfo(buyerInfo: Buyer) {
            buyer = buyerInfo
            bindView()
            nameTextView.text = buyer?.name
            ageTextView.text = buyer?.age.toString()
            idTextView.text = buyer?.id

            dateTextView.text = buyer!!.date.parseDateFromUnixTimestampToDate()
        }


        private fun bindView() {
            nameTextView = itemView.findViewById(R.id.textViewName)
            ageTextView = itemView.findViewById(R.id.textViewAge)
            idTextView = itemView.findViewById(R.id.textViewId)
            dateTextView = itemView.findViewById(R.id.createdAtText)
        }

    }

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): BuyerListViewHolder {
        val itemView = LayoutInflater
            .from(parent.context)
            .inflate(R.layout.buyer_card_view, parent, false)
        return BuyerListViewHolder(itemView)
    }

    override fun onBindViewHolder(holder: BuyerListViewHolder, position: Int) {
        val itemBuyerInfo = buyers[position]
        holder.bindBuyerInfo(itemBuyerInfo)
    }

    override fun getItemCount(): Int {
        return buyers.size
    }

}

