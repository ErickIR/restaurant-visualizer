package com.example.core.domain

data class Transaction(
    val id: String,
    val ipAddress: String,
    val device: String,
    val total: Double,
    val products: ArrayList<Product>
)