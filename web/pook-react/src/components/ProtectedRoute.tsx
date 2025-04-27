import React from "react"
import { Navigate } from "react-router-dom";

export type Props = {
    /** If True, render the child components else redirect to provided url */
    condition: boolean;
    /** The url to redirect to if condition not met */
    redirect: string;
    /** React children */
    children: React.ReactNode;
}

function ProtectedRoute({ condition, redirect, children }: Props) {
    if(!condition) {
        return <Navigate to={redirect} replace/>;
    }
    return <>{children}</>;
}

export default ProtectedRoute;
