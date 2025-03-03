// TODO use generic for data
interface ComponentProp {
    type: string;
    key: string;
    role: string;
    data: { [key: string]: any };
    children: ComponentProp[];
}
