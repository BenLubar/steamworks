// steamtypes.h assumes MSVC on Windows, but we're using GCC.
#define __int16 __INT16_TYPE__
#define __int32 __INT32_TYPE__
#define __int64 __INT64_TYPE__

#define STEAMAPI_WRAPPER
#include "steamworks_wrap.h"
#include <steam/steam_api.h>

#include <map>
#include <memory>
#include <mutex>
#include <vector>

extern "C"
intptr_t SteamAPIWrap_SteamNetworking()
{
    return reinterpret_cast<intptr_t>(SteamNetworking());
}

extern "C"
intptr_t SteamAPIWrap_SteamUser()
{
    return reinterpret_cast<intptr_t>(SteamUser());
}

class CCallbackGo;

static std::recursive_mutex callback_lock;
static std::map<int, std::unique_ptr<CCallbackGo>> callback_lookup;
static std::vector<int> callback_to_delete;

class CCallbackGo : public CCallbackBase
{
public:
    CCallbackGo(size_t size, int callback_type_id, SteamAPICall_t api_call_id = k_uAPICallInvalid) : data_length(size), api_call_id(api_call_id)
    {
        if (api_call_id == k_uAPICallInvalid)
        {
            SteamAPI_RegisterCallback(this, callback_type_id);
        }
        else
        {
            SteamAPI_RegisterCallResult(this, api_call_id);
        }
    }
    ~CCallbackGo()
    {
        if (api_call_id == k_uAPICallInvalid)
        {
            SteamAPI_UnregisterCallback(this);
        }
        else
        {
            SteamAPI_UnregisterCallResult(this, api_call_id);
        }
    }
    virtual void Run(void *data)
    {
        onCallback(GetICallback(), data, data_length, false, api_call_id);

        if (api_call_id != k_uAPICallInvalid)
        {
            callback_to_delete.push_back(GetICallback());
        }
    }
    virtual void Run(void *data, bool ioFailure, SteamAPICall_t hSteamAPICall)
    {
        if (api_call_id != k_uAppIdInvalid && hSteamAPICall != api_call_id)
        {
            return;
        }

        onCallback(GetICallback(), data, data_length, ioFailure, hSteamAPICall);

        if (api_call_id != k_uAPICallInvalid)
        {
            callback_to_delete.push_back(GetICallback());
        }
    }
    virtual int GetCallbackSizeBytes()
    {
        return data_length;
    }

protected:
    const size_t data_length;
    const SteamAPICall_t api_call_id;
};

extern "C"
void SteamAPIWrap_RunCallbacks()
{
    std::lock_guard<std::recursive_mutex> lock(callback_lock);

    SteamAPI_RunCallbacks();

    for (auto id : callback_to_delete)
    {
        callback_lookup.erase(id);
        discardCallback(id);
    }
    callback_to_delete.clear();
}

extern "C"
int SteamAPIWrap_RegisterCallback(size_t payload_size, int callback_id)
{
    std::lock_guard<std::recursive_mutex> lock(callback_lock);

    auto cb = new CCallbackGo(payload_size, callback_id);
    int id = cb->GetICallback();
    callback_lookup[id] = std::unique_ptr<CCallbackGo>(cb);
    return id;
}

extern "C"
int SteamAPIWrap_UnregisterCallback(int callback_id)
{
    std::lock_guard<std::recursive_mutex> lock(callback_lock);

    callback_to_delete.push_back(callback_id);
}

extern "C"
int SteamAPIWrap_RegisterCallResult(size_t payload_size, int callback_id, uint64_t api_call_id)
{
    std::lock_guard<std::recursive_mutex> lock(callback_lock);

    auto cb = new CCallbackGo(payload_size, callback_id, api_call_id);
    int id = cb->GetICallback();
    callback_lookup[id] = std::unique_ptr<CCallbackGo>(cb);
    return id;
}
