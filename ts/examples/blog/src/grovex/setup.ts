import {componentRegistry} from "../grove/Component";
import {Page} from "./Page";
import {Card} from "./Card";
import {Text} from "./Text";

componentRegistry
    .set('grovex.Page', Page)
    .set('grovex.Card', Card)
    .set('grovex.Text', Text);

