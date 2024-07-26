import { createSlice} from "@reduxjs/toolkit";

const initialState = {
    dishes:[]
}

const dishSlice = createSlice({
    name: "dishes",
    initialState,
    reducers: {
        addDish:(state, {payload}) => {
            state.dishes = payload;
        }
    }
});

export const {addDish} = dishSlice.actions
export const getAllDish = (state) => state.dishes.dishes
export default dishSlice.reducer;
