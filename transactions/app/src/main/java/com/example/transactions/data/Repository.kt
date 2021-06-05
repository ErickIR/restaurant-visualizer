package com.example.transactions.data

import com.example.transactions.models.Buyer
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import java.lang.Exception

interface ApiServiceClient {
    companion object {
        private const val baseUrl = "http://192.168.0.102:3000/api/"
        private val retrofit = Retrofit.Builder()
            .baseUrl(baseUrl)
            .addConverterFactory(GsonConverterFactory.create())
            .build()

        @Volatile private var INSTANCE: ApiService? = null
        private val LOCK = Any()

        operator fun invoke(): ApiService = INSTANCE ?: synchronized(LOCK) {
            INSTANCE ?: getApiServiceInstance().also { INSTANCE = it }
        }

        private fun getApiServiceInstance(): ApiService {
            return retrofit.create(ApiService::class.java)
        }
    }
}