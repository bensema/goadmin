import {StrictMode} from "react";
import ReactDOM from "react-dom";
import {BrowserRouter} from "react-router-dom";
import { StyledEngineProvider } from '@mui/material/styles';

import App from "./App";
import {LicenseInfo} from "@mui/x-data-grid-pro";

LicenseInfo.setLicenseKey(process.env.REACT_APP_NEXT_PUBLIC_MUI_LICENSE);

ReactDOM.render(
    <StrictMode>
        <BrowserRouter>
            <StyledEngineProvider injectFirst>
                <App/>
            </StyledEngineProvider>
        </BrowserRouter>
    </StrictMode>,
    document.getElementById("root")
);
