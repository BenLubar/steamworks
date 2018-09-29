#include "shim.h"

#include <map>
#include <memory>
#include <mutex>
#include <steam/steam_api.h>

#include "callback.h"

// Implemented in Go
extern "C" void onCallback(CallbackID_t callback_id, void * data, size_t data_length, bool io_failure, SteamAPICall_t api_call_id);

class CCallbackGo : public CCallbackBase
{
public:
	CCallbackGo(size_t size, int callback_type_id, SteamAPICall_t api_call_id = k_uAPICallInvalid, bool game_server = false) : data_length(size), api_call_id(api_call_id)
	{
		if (game_server)
		{
			m_nCallbackFlags |= k_ECallbackFlagsGameServer;
		}

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
		if (m_nCallbackFlags & k_ECallbackFlagsRegistered)
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
	}
	virtual void Run(void *data)
	{
		onCallback(GetICallback(), data, data_length, false, api_call_id);
	}
	virtual void Run(void *data, bool ioFailure, SteamAPICall_t hSteamAPICall)
	{
		if (api_call_id != k_uAPICallInvalid && hSteamAPICall != api_call_id)
		{
			return;
		}

		onCallback(GetICallback(), data, data_length, ioFailure, hSteamAPICall);
	}
	virtual int GetCallbackSizeBytes()
	{
		return data_length;
	}

protected:
	const size_t data_length;
	const SteamAPICall_t api_call_id;
};

static std::map<CallbackID_t, std::unique_ptr<CCallbackGo>> callbacks;
static std::mutex callbacks_lock;

extern "C" CallbackID_t Register_Callback(size_t size, int callback_type_id, SteamAPICall_t api_call_id, bool game_server)
{
	std::unique_ptr<CCallbackGo> cb(new CCallbackGo(size, callback_type_id, api_call_id, game_server));

	CallbackID_t callback_id = cb->GetICallback();

	{
		std::lock_guard<std::mutex> lock(callbacks_lock);
		callbacks[callback_id] = std::move(cb);
	}

	return callback_id;
}

extern "C" void Unregister_Callback(CallbackID_t callback_id)
{
	std::lock_guard<std::mutex> lock(callbacks_lock);
	callbacks.erase(callback_id);
}

// Implemented in Go
extern "C" void warningMessageHook(int, const char *);

extern "C" void SetWarningMessageHookGo()
{
	SteamUtils()->SetWarningMessageHook(&warningMessageHook);
}
