package com.example.transactions.data

import com.example.transactions.models.Buyer
import okhttp3.OkHttpClient
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import java.lang.Exception
import java.util.concurrent.TimeUnit

interface ApiServiceClient {
    companion object {
        private const val baseUrl = "http://192.168.0.102:3000/api/"
        private val httpClient: OkHttpClient = OkHttpClient.Builder()
            .readTimeout(60, TimeUnit.SECONDS)
            .connectTimeout(60, TimeUnit.SECONDS)
            .build()
        private val retrofit = Retrofit.Builder()
            .baseUrl(baseUrl)
            .addConverterFactory(GsonConverterFactory.create())
            .client(httpClient)
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