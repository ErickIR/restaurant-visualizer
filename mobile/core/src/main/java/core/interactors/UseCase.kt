package com.example.core.interactors

interface UseCase<T> {
    fun execute(): T
}