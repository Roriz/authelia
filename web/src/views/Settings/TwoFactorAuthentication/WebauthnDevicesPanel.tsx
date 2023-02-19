import React, { Fragment, Suspense, useState } from "react";

import { Box, Button, Paper, Stack, Tooltip, Typography } from "@mui/material";
import { useTranslation } from "react-i18next";

import { AutheliaState } from "@services/State";
import LoadingPage from "@views/LoadingPage/LoadingPage";
import WebauthnDeviceRegisterDialog from "@views/Settings/TwoFactorAuthentication/WebauthnDeviceRegisterDialog";
import WebauthnDevicesStack from "@views/Settings/TwoFactorAuthentication/WebauthnDevicesStack";

interface Props {
    state: AutheliaState;
}

export default function WebauthnDevicesPanel(props: Props) {
    const { t: translate } = useTranslation("settings");

    const [showWebauthnDeviceRegisterDialog, setShowWebauthnDeviceRegisterDialog] = useState<boolean>(false);
    const [refreshState, setRefreshState] = useState<number>(0);

    const handleIncrementRefreshState = () => {
        setRefreshState((refreshState) => refreshState + 1);
    };

    return (
        <Fragment>
            <WebauthnDeviceRegisterDialog
                open={showWebauthnDeviceRegisterDialog}
                onClose={() => {
                    handleIncrementRefreshState();
                }}
                setCancelled={() => {
                    setShowWebauthnDeviceRegisterDialog(false);
                    handleIncrementRefreshState();
                }}
            />
            <Paper variant="outlined">
                <Box sx={{ p: 3 }}>
                    <Stack spacing={2}>
                        <Box>
                            <Typography variant="h5">{translate("Webauthn Credentials")}</Typography>
                        </Box>
                        <Box>
                            <Tooltip title={translate("Click to add a Webauthn credential to your account")}>
                                <Button
                                    variant="outlined"
                                    color="primary"
                                    onClick={() => {
                                        setShowWebauthnDeviceRegisterDialog(true);
                                    }}
                                >
                                    {translate("Add Credential")}
                                </Button>
                            </Tooltip>
                        </Box>
                        <Suspense fallback={<LoadingPage />}>
                            <WebauthnDevicesStack
                                refreshState={refreshState}
                                incrementRefreshState={handleIncrementRefreshState}
                            />
                        </Suspense>
                    </Stack>
                </Box>
            </Paper>
        </Fragment>
    );
}
