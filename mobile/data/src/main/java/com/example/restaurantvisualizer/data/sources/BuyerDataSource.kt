package com.example.restaurantvisualizer.data.sources

import com.example.resturantvisualizer.domain.models.Buyer

interface BuyerDataSource {
    suspend fun getAllBuyers(page: Int, size: Int): List<Buyer>
}