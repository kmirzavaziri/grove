import {componentRegistry} from "../grove/Component";
import {DTypography} from "./DTypography";
import {LPage} from "./LPage";
import {DProfile} from "./DProfile";
import {DDivider} from "./DDivider";
import {DBreadcrumbs} from "./DBreadcrumbs";
import {XClickable} from "./XClickable";
import {DButton} from "./DButton";
import {LBox} from "./LBox";
import {LFragment} from "./LFragment";

componentRegistry
    .set('grovex.DTypography', DTypography)
    .set('grovex.LPage', LPage)
    .set('grovex.LBox', LBox)
    .set('grovex.LFragment', LFragment)
    .set('grovex.DProfile', DProfile)
    .set('grovex.DDivider', DDivider)
    .set('grovex.DBreadcrumbs', DBreadcrumbs)
    .set('grovex.DButton', DButton)
    .set('grovex.XClickable', XClickable)

