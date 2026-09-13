import { create } from "zustand";
import type { User } from "../models/user"


interface UserState {
    user: User | null
    setUser: (newUser: User) => void
}

export const useUser = create<UserState>((set) => ({
    user: null,
    setUser: (newUser: User) => set({ user: newUser })
}))