package com.example.transactions.presentation.buyerdetails.adapters

import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.recyclerview.widget.RecyclerView
import com.example.transactions.R
import com.example.transactions.models.Product

class ProductstAdapter(private val products: ArrayList<Product>)
    : RecyclerView.Adapter<ProductstAdapter.ProductsViewHolder>() {
    class ProductsViewHolder(itemView: View) : RecyclerView.ViewHolder(itemView) {
        private var product: Product? = null
        private lateinit var nameTextView: TextView
        private lateinit var priceTextView: TextView

        fun bindProductInfo(productInfo: Product) {
            product = productInfo
            bindView()
            nameTextView.text = if(productInfo.name.length >= 35) "${productInfo.name.substring(0..35)}..." else productInfo.name
            priceTextView.text = "Price: ${productInfo.price}"
        }

        private fun bindView() {
            nameTextView = itemView.findViewById(R.id.productNameTextView)
            priceTextView = itemView.findViewById(R.id.priceTextView)
        }
    }

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): ProductsViewHolder {
        val itemView = LayoutInflater
            .from(parent.context)
            .inflate(R.layout.product_item, parent, false)
        return ProductsViewHolder(itemView)
    }

    override fun onBindViewHolder(holder: ProductsViewHolder, position: Int) {
        val item = products[position]
        holder.bindProductInfo(item)
    }

    override fun getItemCount(): Int {
        return products.size
    }
}