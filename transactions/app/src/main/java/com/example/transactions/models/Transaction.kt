package com.example.transactions.models

import com.google.gson.annotations.SerializedName

data class Transaction(
    @SerializedName("id")
    val id: String,
    @SerializedName("ipAddress")
    val ipAddress: String,
    @SerializedName("device")
    val device: String,
    @SerializedName("total")
    val total: Int,
    @SerializedName("products")
    val products: List<Product>
)
