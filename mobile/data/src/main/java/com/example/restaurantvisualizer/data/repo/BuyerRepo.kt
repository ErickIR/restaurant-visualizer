package com.example.restaurantvisualizer.data.repo

import com.example.restaurantvisualizer.data.sources.BuyerDataSource
import com.example.resturantvisualizer.domain.models.Buyer

class BuyerRepo(private val dataSource: BuyerDataSource)  {
    suspend fun getAllBuyers(page: Int, size: Int): List<Buyer> {
        return dataSource.getAllBuyers(page, size)
    }
}