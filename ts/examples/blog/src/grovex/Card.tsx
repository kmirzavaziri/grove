import React from "react";
import {Component, ComponentProps} from "../grove/Component";

export interface CardProps {
    title: string
    children: ComponentProps[]
}

export const Card: React.FC<CardProps> = (props) => (
    <div className="card">
        {props.title && <h2>{props.title}</h2>}
        {props.children.map(child => <Component key={child.key} props={child}/>)}
    </div>
);
