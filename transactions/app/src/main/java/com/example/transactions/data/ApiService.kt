package com.example.transactions.data

import retrofit2.Call
import retrofit2.http.GET
import retrofit2.http.POST
import retrofit2.http.Query

interface ApiService {
    @GET("/api/buyer")
    fun getAllBuyers(@Query("page") page: Int) : Call<PaginatedApiResponse>

    @POST("/api/load")
    fun loadData(@Query("date") date: String) : Call<ApiResponse<Unit>>
}