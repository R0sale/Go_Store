import * as React from 'react'
import ErrorBoundary from './ErrorBoundary/ErrorBoundary'
import { Header } from './Header'

export const Layout = ({ children }: { children: React.ReactNode }) => {
    return (
        <ErrorBoundary fallback={<p>Sorry</p>}>
            <div className="min-h-screen bg-slate-50">
                <Header />
                <ErrorBoundary fallback={<p>Sorry, cant load anything</p>}>
                    {children}
                </ErrorBoundary>
            </div>
        </ErrorBoundary>
    )
}