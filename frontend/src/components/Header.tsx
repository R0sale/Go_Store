import { useUser } from "../stores/user_store"
import guest from "../assets/guest.png"
import { Package, ShoppingBasket } from "lucide-react"
import { useNavigate } from "react-router"

export const Header = () => {
    const user = useUser(state => state.user)
    const navigate = useNavigate()

    const goToCart = () => {
        navigate("/cart")
    }

    const goToMain = () => {
        navigate("/")
    }

    const goToOrders = () => {
        navigate("/orders")
    }

    const goToProfile = () => {
        navigate("/profile")
    }

    return (
        <header className="flex items-center justify-between px-6 py-4 bg-white shadow-sm">
            <div className="flex items-center gap-3">
                <header onClick={goToMain} className="text-xl font-bold text-gray-900 cursor-pointer">Go Store</header>
            </div>
            <div className="flex items-center gap-6">
                <Package onClick={goToCart} className="cursor-pointer" />
                <ShoppingBasket onClick={goToOrders} className="cursor-pointer" />
                <img src={user?.ImageUrl == "" ? guest : user?.ImageUrl} className="w-9 h-9 rounded-full object-cover cursor-pointer" onClick={goToProfile} />
            </div>
        </header>
    )
}