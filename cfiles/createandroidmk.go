package cfiles

import (
	"fmt"
	"log"
	"os"
)

func (ctx *Ctactx) WriteAndroid() {
	android := `###################### optee-hello-world ######################
LOCAL_PATH := $(call my-dir)

include $(CLEAR_VARS)
LOCAL_CFLAGS += -DANDROID_BUILD
LOCAL_CFLAGS += -Wall

LOCAL_SRC_FILES += host/main.c

LOCAL_C_INCLUDES := $(LOCAL_PATH)/ta/include

LOCAL_SHARED_LIBRARIES := libteec
LOCAL_MODULE := %s
LOCAL_VENDOR_MODULE := true
LOCAL_MODULE_TAGS := optional
include $(BUILD_EXECUTABLE)

include $(LOCAL_PATH)/ta/Android.mk
`
	err := os.WriteFile("./Android.mk", []byte(fmt.Sprintf(android, ctx.lowerName)), MODE)
	if err != nil {
		log.Print("Error creating the Android.mk file: ", err.Error())
	}
}
