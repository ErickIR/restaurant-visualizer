package com.example.transactions.presentation.buyerlist

import android.util.Log
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.recyclerview.widget.RecyclerView
import com.example.transactions.R
import com.example.transactions.models.Buyer

class BuyerListAdapter(private val buyers: ArrayList<Buyer>)
    : RecyclerView.Adapter<BuyerListAdapter.BuyerListViewHolder>() {

    class BuyerListViewHolder(itemView: View) : RecyclerView.ViewHolder(itemView), View.OnClickListener {
        private var buyer: Buyer? = null
        private lateinit var nameTextView: TextView
        private lateinit var ageTextView: TextView
        private lateinit var idTextView: TextView

        init {
            itemView.setOnClickListener(this)
        }

        override fun onClick(p0: View?) {
            Log.d("RecyclerView", "${buyer?.name} with age: ${buyer?.age} was clicked!")
        }

        fun bindBuyerInfo(buyerInfo: Buyer) {
            buyer = buyerInfo
            bindView()
            nameTextView.text = buyer?.name
            ageTextView.text = buyer?.age.toString()
            idTextView.text = buyer?.id
        }

        private fun bindView() {
            nameTextView = itemView.findViewById(R.id.textViewName)
            ageTextView = itemView.findViewById(R.id.textViewAge)
            idTextView = itemView.findViewById(R.id.textViewId)
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