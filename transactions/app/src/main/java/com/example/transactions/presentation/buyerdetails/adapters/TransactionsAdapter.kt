package com.example.transactions.presentation.buyerdetails.adapters

import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.RecyclerView
import com.example.transactions.R
import com.example.transactions.models.Product
import com.example.transactions.models.Transaction

class TransactionsAdapter(private val transactions: ArrayList<Transaction>)
    : RecyclerView.Adapter<TransactionsAdapter.TransactionViewHolder>() {
        class TransactionViewHolder(itemView: View): RecyclerView.ViewHolder(itemView), View.OnClickListener {
            private var transaction: Transaction? = null
            private lateinit var transactionIdTextView: TextView
            private lateinit var totalTextView: TextView
            private lateinit var transactionsRecyclerView: RecyclerView
            private var showProducts = false
            private lateinit var adapter: ProductstAdapter
            private val products = ArrayList<Product>()
            init {
                itemView.setOnClickListener(this)
            }

            fun bindTransactionInfo(transactionInfo: Transaction) {
                transaction = transactionInfo
                bindView()
                transactionIdTextView.text = transactionInfo.id
                totalTextView.text = "Total: ${transactionInfo.total}"
                if (transactionInfo.products != null) {
                    products.addAll(transactionInfo.products)
                }

                adapter = ProductstAdapter(products)
                transactionsRecyclerView.layoutManager = LinearLayoutManager(itemView.context)
                transactionsRecyclerView.adapter = adapter
            }

            private fun bindView() {
                transactionIdTextView = itemView.findViewById(R.id.transactionIdTextView)
                totalTextView = itemView.findViewById(R.id.totalTextView)
                transactionsRecyclerView = itemView.findViewById(R.id.transactionsRecyclerView)
            }

            override fun onClick(view: View?) {
                showProducts = !showProducts
                transactionsRecyclerView.visibility = if (showProducts) View.VISIBLE else View.GONE
            }

        }

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): TransactionViewHolder {
        val itemView = LayoutInflater
            .from(parent.context)
            .inflate(R.layout.transaction_items, parent, false)
        return TransactionViewHolder(itemView)
    }

    override fun onBindViewHolder(holder: TransactionViewHolder, position: Int) {
        val itemInfo = transactions[position]
        holder.bindTransactionInfo(itemInfo)
    }

    override fun getItemCount(): Int {
        return transactions.size
    }
}