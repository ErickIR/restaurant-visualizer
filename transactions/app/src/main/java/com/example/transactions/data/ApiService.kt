package com.example.transactions.data

import com.example.transactions.data.dto.BuyerDetailsDto
import com.example.transactions.models.Buyer
import retrofit2.Call
import retrofit2.http.GET
import retrofit2.http.POST
import retrofit2.http.Path
import retrofit2.http.Query

interface ApiService {
    @GET("/api/buyer/all")
    fun getBuyersByDate(@Query("date") date: String) : Call<ApiResponse<List<Buyer>>>

    @GET("/api/buyer")
    fun getAllBuyers(@Query("page") page: Int) : Call<PaginatedApiResponse>

    @GET("/api/buyer/{buyerId}")
    fun getBuyerInformation(@Path("buyerId") buyerId: String): Call<ApiResponse<BuyerDetailsDto>>

    @POST("/api/load")
    fun loadData(@Query("date") date: String) : Call<ApiResponse<Unit>>
}