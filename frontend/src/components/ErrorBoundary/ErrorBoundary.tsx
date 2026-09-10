import * as React from 'react'

interface ErrorBoundaryProps {
    fallback: React.ReactNode
    children: React.ReactNode
}

interface State {
    hasError: boolean
}

class ErrorBoundary extends React.Component<ErrorBoundaryProps, State> {
    public state: State = {
        hasError: false
    }

    public static getDerivedStateFromError(error: Error): State {
        return { hasError: true }
    }

    componentDidCatch(error: Error, info: React.ErrorInfo) {
        console.error("Uncaught error:", error, info)
    }

    render() {
        if (this.state.hasError) {
            return this.props.fallback
        }

        return this.props.children
    }
}

export default ErrorBoundary