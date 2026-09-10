

export const Fallback = () => {
    return (
        <div className="flex flex-col items-center justify-center min-h-screen gap-3 px-6 text-center bg-white">
            <header className="text-xl font-bold text-gray-900">Go Store</header>
            <div className="text-gray-600">Sorry, we couldn't load the page. Click refresh button to try again</div>
        </div>
    )
}