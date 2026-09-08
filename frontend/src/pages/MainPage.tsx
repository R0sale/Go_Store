import { useLoaderData } from "react-router"
import type { item } from "../models/item"
import { useState } from "react"
import { Catalog } from "../components/Catalog"


export const loadCatalogItems = async (): Promise<item[] | null> => {
    const response = await fetch("", {
        method: "GET",
        credentials: "include"
    })

    if (!response.ok) {
        return null
    }

    const result = await response.json() as item[]

    return result
}

export const MainPage = () => {
    const items = useLoaderData<item[] | null>()
    const [catalogItems, setCatalogItems] = useState<item[] | null>(items)

    return (
        <div>
            <Catalog items={catalogItems}/>
        </div>
    )
}