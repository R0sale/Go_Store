import type { item } from "../models/item"

interface CatalogProps {
    items: item[] | null
}

export const Catalog = (props: CatalogProps) => {
    return (
        <div className="mx-auto grid max-w-7xl grid-cols-2 gap-4 p-4 sm:grid-cols-3 sm:gap-6 sm:p-6 lg:grid-cols-4 xl:grid-cols-5">
            {props.items?.map(item =>
                <div
                    key={item?.Id}
                    className="group flex flex-col overflow-hidden rounded-2xl border border-slate-200 bg-white shadow-sm transition duration-300 hover:-translate-y-1 hover:shadow-xl"
                >
                    <div className="aspect-square w-full overflow-hidden bg-slate-100">
                        <img
                            src={item?.ImageUrl}
                            alt="item image"
                            className="h-full w-full object-cover transition duration-300 group-hover:scale-105"
                        />
                    </div>
                    <div className="flex flex-1 flex-col gap-1 p-4">
                        <p className="truncate text-sm font-medium text-slate-800">{item?.Name}</p>
                        <p className="mt-auto text-lg font-bold text-indigo-600">${item?.Price.toFixed(2)}</p>
                    </div>
                </div>
            )}
        </div>
    )
}