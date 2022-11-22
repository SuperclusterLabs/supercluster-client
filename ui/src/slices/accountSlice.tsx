import { createSlice } from "@reduxjs/toolkit"
import type { PayloadAction } from "@reduxjs/toolkit"
import type { RootState } from "../store/store"

interface AccountState {
  address: string
}

const initialState: AccountState = {
  address: "",
}

export const accountSlice = createSlice({
  name: 'account',
  initialState,
  reducers: {
    setAccount: (state, action: PayloadAction<string>) => {
      state.address = action.payload
    }
  }
})

export const { setAccount } = accountSlice.actions

export const selectAccount = (state: RootState) => state.account.address

export default accountSlice.reducer
