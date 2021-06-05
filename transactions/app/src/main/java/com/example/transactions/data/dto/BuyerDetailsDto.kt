package com.example.transactions.models

import com.example.transactions.data.dto.BuyersWithSameIp
import com.google.gson.annotations.SerializedName

data class BuyerDetailsDto(
    @SerializedName("buyer")
    val buyer: Buyer,
    @SerializedName("transactions")
    val transactions: List<Transaction>,
    @SerializedName("buyerWithSameIp")
    val buyersWithSameIp: List<BuyersWithSameIp>
)
