import { useLoaderData } from "react-router"
import type { item } from "../models/item"
import { useState } from "react"
import { Catalog } from "../components/Catalog"
import { Header } from "../components/Header"


export const loadCatalogItems = async (): Promise<item[] | null> => {
    console.log(import.meta.env.VITE_BASE_URL + "/api/catalog")

    const response = await fetch(import.meta.env.VITE_BASE_URL + "/api/catalog", {
        method: "GET",
        credentials: "include"
    })

    if (!response.ok) {
        return null
    }

    const result = await response.json() as item[]

    console.log(result)

    return result
}

export const MainPage = () => {
    const items = useLoaderData<item[] | null>()
    const [catalogItems, setCatalogItems] = useState<item[] | null>(items)

    return (
        <div className="min-h-screen bg-slate-50">
            <Header />
            <Catalog items={catalogItems}/>
        </div>
    )
}