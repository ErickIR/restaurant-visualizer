package com.example.transactions.data.dto

import com.example.transactions.models.Buyer
import com.example.transactions.models.Product
import com.example.transactions.models.Transaction
import com.google.gson.annotations.SerializedName

data class BuyerDetailsDto(
    @SerializedName("buyer")
    val buyer: Buyer,
    @SerializedName("transactions")
    val transactions: List<Transaction>,
    @SerializedName("buyerWithSameIp")
    val buyersWithSameIp: List<BuyersWithSameIp>,
    @SerializedName("products")
    val products: List<Product>
)
