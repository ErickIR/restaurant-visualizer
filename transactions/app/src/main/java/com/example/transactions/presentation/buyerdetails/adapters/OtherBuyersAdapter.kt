package com.example.transactions.presentation.buyerdetails.adapters

import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.recyclerview.widget.RecyclerView
import com.example.transactions.R
import com.example.transactions.models.Buyer

class OtherBuyersAdapter(private val otherBuyers: ArrayList<Buyer>)
    : RecyclerView.Adapter<OtherBuyersAdapter.OtherBuyersViewHolder>() {

    class OtherBuyersViewHolder(itemView: View) : RecyclerView.ViewHolder(itemView) {
        private var buyer: Buyer? = null
        private lateinit var nameTextView: TextView
        private lateinit var ageTextView: TextView
        private lateinit var idTextView: TextView

        fun bindBuyerInfo(buyerInfo: Buyer) {
            buyer = buyerInfo
            bindView()
            nameTextView.text = buyerInfo.name
            ageTextView.text = "Age: ${buyerInfo.age}"
            idTextView.text = "ID: ${buyerInfo.id}"
        }

        private fun bindView() {
            nameTextView = itemView.findViewById(R.id.itemTitle)
            ageTextView = itemView.findViewById(R.id.itemAge)
            idTextView = itemView.findViewById(R.id.itemId)
        }

    }

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): OtherBuyersViewHolder {
        val itemView = LayoutInflater
            .from(parent.context)
            .inflate(R.layout.other_buyers_item, parent, false)
        return OtherBuyersViewHolder(itemView)
    }

    override fun onBindViewHolder(holder: OtherBuyersViewHolder, position: Int) {
        val itemBuyerInfo = otherBuyers[position]
        holder.bindBuyerInfo(itemBuyerInfo)
    }

    override fun getItemCount(): Int {
        return otherBuyers.size
    }
}