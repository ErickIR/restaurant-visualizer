package com.example.transactions.presentation.helper

import java.text.SimpleDateFormat
import java.util.*

fun String.parseDateFromUnixTimestampToDate(): String {
    val dateFormat = SimpleDateFormat("yyyy-MM-dd")
    val date = Date(this.toLong() * 1000)
    return dateFormat.format(date)
}