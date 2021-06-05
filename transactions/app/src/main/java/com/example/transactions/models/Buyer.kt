package com.example.transactions.models

import com.google.gson.annotations.SerializedName

data class Buyer(
        @SerializedName("id")
        val id: String,
        @SerializedName("name")
        val name: String,
        @SerializedName("age")
        val age: Int,
        @SerializedName("date")
        val date: String
)
