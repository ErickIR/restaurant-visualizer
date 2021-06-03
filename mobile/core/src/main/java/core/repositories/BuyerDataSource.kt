package com.example.core.repositories

import com.example.core.domain.Buyer

interface BuyerDataSource {
    fun getAllBuyers(): List<Buyer>
}