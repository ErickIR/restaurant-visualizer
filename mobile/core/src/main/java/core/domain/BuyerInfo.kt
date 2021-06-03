package com.example.core.domain

data class BuyerInfo (
    val buyer: Buyer,
    val transactions: List<Transaction>,
    val buyersWithSameIp: List<BuyersWithRelatedIps>
)

data class BuyersWithRelatedIps(
    val buyer: Buyer,
    val ipAddress: String,
    val device: String
)