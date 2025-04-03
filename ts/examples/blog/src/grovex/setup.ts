import {componentRegistry} from "../grove/Component";
import {DBreadcrumbs} from "./DBreadcrumbs";
import {DButton} from "./DButton";
import {DDivider} from "./DDivider";
// import {DImage} from "./DImage";
import {DProfile} from "./DProfile";
import {DTypography} from "./DTypography";
// import {IBool} from "./IBool";
// import {IDatetime} from "./IDatetime";
// import {IList} from "./IList";
// import {IMultiSelect} from "./IMultiSelect";
// import {INumber} from "./INumber";
// import {IPhoneNumber} from "./IPhoneNumber";
// import {ISingleSelect} from "./ISingleSelect";
import {IText} from "./IText";
// import {LAccordion} from "./LAccordion";
import {LBox} from "./LBox";
import {LFragment} from "./LFragment";
// import {LGrid} from "./LGrid";
// import {LMasonry} from "./LMasonry";
import {LPage} from "./LPage";
// import {LTabs} from "./LTabs";
import {XClickable} from "./XClickable";
import {XModal} from "./XModal";

import {actionRegistry} from "../grove/action";
import {ARender} from "./ARender";

componentRegistry
    .set('grovex.DTypography', DTypography)
    .set('grovex.DBreadcrumbs', DBreadcrumbs)
    .set('grovex.DButton', DButton)
    .set('grovex.DDivider', DDivider)
    // .set('grovex.DImage', DImage)
    .set('grovex.DProfile', DProfile)
    .set('grovex.DTypography', DTypography)
    // .set('grovex.IBool', IBool)
    // .set('grovex.IDatetime', IDatetime)
    // .set('grovex.IList', IList)
    // .set('grovex.IMultiSelect', IMultiSelect)
    // .set('grovex.INumber', INumber)
    // .set('grovex.IPhoneNumber', IPhoneNumber)
    // .set('grovex.ISingleSelect', ISingleSelect)
    .set('grovex.IText', IText)
    // .set('grovex.LAccordion', LAccordion)
    .set('grovex.LBox', LBox)
    .set('grovex.LFragment', LFragment)
    // .set('grovex.LGrid', LGrid)
    // .set('grovex.LMasonry', LMasonry)
    .set('grovex.LPage', LPage)
    // .set('grovex.LTabs', LTabs)
    .set('grovex.XClickable', XClickable)
    .set('grovex.XModal', XModal)

actionRegistry
    .set("grovex.ARender", ARender)