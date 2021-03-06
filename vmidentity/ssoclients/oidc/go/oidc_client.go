package oidc

/*
#cgo CFLAGS: -I/root/git/lightwave/vmidentity/ssoclients/common/include/public/
#cgo CFLAGS: -I/root/git/lightwave/vmidentity/ssoclients/oidc/include/public/
#cgo LDFLAGS: -L/root/git/lightwave/vmidentity/build/ssoclients/common/src/.libs/ -l ssocommon
#cgo LDFLAGS: -L/root/git/lightwave/vmidentity/build/ssoclients/oidc/src/.libs/ -l ssooidc
#include <stdlib.h>
#include "ssotypes.h"
#include "ssoerrors.h"
#include "oidc_types.h"
#include "oidc.h"
*/
import "C"

import "runtime"

type OidcClient struct {
    p C.POIDC_CLIENT
}

func OidcClientGlobalInit() (err error) {
    var e C.SSOERROR = C.OidcClientGlobalInit()
    if e != 0 {
        err = cErrorToGoError(e)
    }
    return err
}

func OidcClientGlobalCleanup() {
    C.OidcClientGlobalCleanup()
}

func OidcClientBuild(
        server string,
        portNumber int,
        tenant string) (result *OidcClient, err error) {
    serverCStr := goStringToCString(server)
    tenantCStr := goStringToCString(tenant)

    defer freeCString(serverCStr)
    defer freeCString(tenantCStr)

    var p C.POIDC_CLIENT = nil
    var e C.SSOERROR = C.OidcClientBuild(
        &p,
        serverCStr,
        C.int(portNumber),
        tenantCStr)
    if e != 0 {
        err = cErrorToGoError(e)
        return
    }

    result = &OidcClient { p }
    runtime.SetFinalizer(result, oidcClientFinalize)
    return
}

func oidcClientFinalize(this *OidcClient) {
    this.Close()
}

func (this *OidcClient) Close() error {
    if (this.p != nil) {
        C.OidcClientDelete(this.p)
        this.p = nil
    }
    return nil
}

func (this *OidcClient) AcquireTokensByPassword(
        username string,
        password string,
        scope string) (successResponse *TokenSuccessResponse, errorResponse *ErrorResponse, err error) {
    usernameCStr    := goStringToCString(username)
    passwordCStr    := goStringToCString(password)
    scopeCStr       := goStringToCString(scope)

    defer freeCString(usernameCStr)
    defer freeCString(passwordCStr)
    defer freeCString(scopeCStr)

    var oidcTokenSuccessResponse C.POIDC_TOKEN_SUCCESS_RESPONSE = nil
    var oidcErrorResponse C.POIDC_ERROR_RESPONSE = nil
    var e C.SSOERROR = C.OidcClientAcquireTokensByPassword(
        this.p,
        usernameCStr,
        passwordCStr,
        scopeCStr,
        &oidcTokenSuccessResponse,
        &oidcErrorResponse)
    if e != 0 {
        err = cErrorToGoError(e)
        return
    }

    if oidcTokenSuccessResponse != nil {
        successResponse = tokenSuccessResponseNew(oidcTokenSuccessResponse)
    }
    if oidcErrorResponse != nil {
        errorResponse = errorResponseNew(oidcErrorResponse)
    }
    return
}

func (this *OidcClient) AcquireTokensByRefreshToken(
        refreshToken string) (successResponse *TokenSuccessResponse, errorResponse *ErrorResponse, err error) {
    refreshTokenCStr := goStringToCString(refreshToken)
    defer freeCString(refreshTokenCStr)

    var oidcTokenSuccessResponse C.POIDC_TOKEN_SUCCESS_RESPONSE = nil
    var oidcErrorResponse C.POIDC_ERROR_RESPONSE = nil
    var e C.SSOERROR = C.OidcClientAcquireTokensByRefreshToken(
        this.p,
        refreshTokenCStr,
        &oidcTokenSuccessResponse,
        &oidcErrorResponse)
    if e != 0 {
        err = cErrorToGoError(e)
        return
    }

    if oidcTokenSuccessResponse != nil {
        successResponse = tokenSuccessResponseNew(oidcTokenSuccessResponse)
    }
    if oidcErrorResponse != nil {
        errorResponse = errorResponseNew(oidcErrorResponse)
    }
    return
}

func (this *OidcClient) AcquireTokensBySolutionUserCredentials(
        certificateSubjectDN string,
        privateKeyPEM string,
        scope string) (successResponse *TokenSuccessResponse, errorResponse *ErrorResponse, err error) {
    certificateSubjectDNCStr    := goStringToCString(certificateSubjectDN)
    privateKeyPEMCStr           := goStringToCString(privateKeyPEM)
    scopeCStr                   := goStringToCString(scope)

    defer freeCString(certificateSubjectDNCStr)
    defer freeCString(privateKeyPEMCStr)
    defer freeCString(scopeCStr)

    var oidcTokenSuccessResponse C.POIDC_TOKEN_SUCCESS_RESPONSE = nil
    var oidcErrorResponse C.POIDC_ERROR_RESPONSE = nil
    var e C.SSOERROR = C.OidcClientAcquireTokensBySolutionUserCredentials(
        this.p,
        certificateSubjectDNCStr,
        privateKeyPEMCStr,
        scopeCStr,
        &oidcTokenSuccessResponse,
        &oidcErrorResponse)
    if e != 0 {
        err = cErrorToGoError(e)
        return
    }

    if oidcTokenSuccessResponse != nil {
        successResponse = tokenSuccessResponseNew(oidcTokenSuccessResponse)
    }
    if oidcErrorResponse != nil {
        errorResponse = errorResponseNew(oidcErrorResponse)
    }
    return
}

func (this *OidcClient) GetSigningCertificatePEM() string {
    return cStringToGoString(C.OidcClientGetSigningCertificatePEM(this.p))
}
