import React from "react";
import {Component, ComponentProps} from "../grove/Component";
import {Action, perform} from "./action";

export interface XClickable {
    action: Action;
    children: ComponentProps[];
}

export const XClickable: React.FC<XClickable> = (props) => (
    // TODO use MUI?
    <div onClick={() => perform(props.action)}>
        {props.children.map(child => (<Component key={child.key} props={child}/>))}
    </div>
);
