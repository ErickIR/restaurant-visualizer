package com.example.transactions.data.dto

import com.example.transactions.models.Buyer
import com.google.gson.annotations.SerializedName

data class BuyersWithSameIp(
    @SerializedName("device")
    val device: String,
    @SerializedName("ipAddress")
    val ipAddress: String,
    @SerializedName("buyer")
    val buyer: Buyer
)
