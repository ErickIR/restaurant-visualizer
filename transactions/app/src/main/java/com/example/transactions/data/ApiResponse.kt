package com.example.transactions.data

import com.example.transactions.models.Buyer
import com.google.gson.annotations.SerializedName

class ApiResponse<T>(
    @SerializedName("data")
    val data: T,
    @SerializedName("message")
    val message: String,
    @SerializedName("success")
    val success: Boolean,
)

data class PaginatedApiResponse(
    @SerializedName("data")
    val data: List<Buyer>,
    @SerializedName("message")
    val message: String,
    @SerializedName("success")
    val success: Boolean,
    @SerializedName("metadata")
    val metadata: Metadata
)

data class Metadata(
    @SerializedName("page")
    val page: Int,
    @SerializedName("size")
    val size: Int,
    @SerializedName("totalPages")
    val totalPages: Int,
    @SerializedName("totalSize")
    val totalSize: Int,
    @SerializedName("next")
    val next: String,
    @SerializedName("previous")
    val previous: String,
)