#pragma once

#ifdef __cplusplus
extern "C"
{
#endif

typedef int CallbackID_t;
extern CallbackID_t Register_Callback(size_t size, int callback_type_id, SteamAPICall_t api_call_id, bool game_server);
extern void Unregister_Callback(CallbackID_t callback_id);
extern void SetWarningMessageHookGo();

#ifdef __cplusplus
}
#endif
