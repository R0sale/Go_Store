import { EllipsisVertical } from "lucide-react"

export const Header = () => {


    return (
        <header className="flex items-center justify-between px-6 py-4 bg-white shadow-sm">
            <div className="flex items-center gap-3">
                <EllipsisVertical className="w-6 h-6 cursor-pointer text-gray-700 hover:text-gray-900" />
                <header className="text-xl font-bold text-gray-900">Go Store</header>
            </div>

            <img className="w-9 h-9 rounded-full object-cover" />
        </header>
    )
}