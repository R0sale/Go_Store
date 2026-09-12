import type { User } from "../models/user"
import { newApiService } from "../services/api_service"
import { useState } from "react"
import { useNavigate } from "react-router"
import { useUser } from "../stores/user_store"

export const LoginPage = () => {
    const [isValid, setIsValid] = useState<boolean>(true)
    const [email, setEmail] = useState<string>("")
    const [password, setPassword] = useState<string>("")
    const navigate = useNavigate()
    const setUser = useUser(state => state.setUser)

    const inputEmail = (e: React.ChangeEvent<HTMLInputElement>) => {
        setEmail(e.target.value)
    }

    const inputPassword = (e: React.ChangeEvent<HTMLInputElement>) => {
        setPassword(e.target.value)
    }

    const login = async (e: React.SubmitEvent) => {
        e.preventDefault()
        const apiService = newApiService()

        const body = {
            email: email,
            password: password
        }

        console.log(import.meta.env.VITE_USERS_URL)
        const { ok, response } = await apiService.apiCall(import.meta.env.VITE_USERS_URL + "/api/users/login", "POST", body)

        if (!ok) {
            setIsValid(false)
            return
        }

        const user = response as User
        setUser(user)

        navigate("/")
    }

    return (
        <div className="mx-auto my-20 flex w-full max-w-sm flex-col gap-5 rounded-2xl border border-gray-100 bg-white p-8 shadow-xl">
            <h1 className="font-serif text-3xl font-bold text-gray-900">Welcome back</h1>
            <p className="-mt-3 text-sm text-gray-500">Don't have an account?<span className="ml-1 cursor-pointer font-medium text-indigo-600 hover:underline">Sign up free</span></p>
            <div className="flex cursor-pointer items-center justify-center gap-2 rounded-lg border border-gray-200 bg-gray-50 py-2.5 text-sm font-medium text-gray-700 hover:bg-gray-100">Google</div>
            <div className="relative my-1 text-center text-xs text-gray-400 before:absolute before:top-1/2 before:left-0 before:h-px before:w-[35%] before:bg-gray-200 after:absolute after:top-1/2 after:right-0 after:h-px after:w-[35%] after:bg-gray-200">or email</div>
            <form onSubmit={login} className="flex flex-col gap-5">
                <label htmlFor="email" className="-mb-3 block text-sm font-medium text-gray-700">Email</label>
                <input onChange={(e) => inputEmail(e)} name="email" className={`rounded-lg border ${!isValid && "border-red-400 border-2"} border-gray-200 bg-gray-50 px-3 py-2.5 text-sm text-gray-900 placeholder:text-gray-400 focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500`} />
                <label htmlFor="password" className="-mb-3 block text-sm font-medium text-gray-700">Password</label>
                <input onChange={(e) => inputPassword(e)} name="password" className={`rounded-lg border ${!isValid && "border-red-400 border-2"} border-gray-200 bg-gray-50 px-3 py-2.5 text-sm text-gray-900 placeholder:text-gray-400 focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500`} />

                {!isValid && <p className="-mb-3 text-red-400 mt-0 pt-0 block text-sm font-medium">Wrong email or password</p>}

                <button type="submit" className="mt-1 rounded-lg bg-indigo-600 py-2.5 text-sm font-semibold text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">Sign in</button>
            </form>
            <footer className="text-center text-xs text-gray-400">By signing in you agree to our Terms and Privacy Policy</footer>
        </div>
    )
}