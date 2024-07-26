import { configureStore } from "@reduxjs/toolkit";
import dishReducer from './dishSlice'

export const store = configureStore({
    reducer: {
        dishes: dishReducer
    }
})