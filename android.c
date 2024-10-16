//go:build android

#include <jni.h>
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>

static jclass find_class(JNIEnv *env, const char *class_name) {
	jclass clazz = (*env)->FindClass(env, class_name);
	if (clazz == NULL) {
		(*env)->ExceptionClear(env);
		printf("cannot find %s", class_name);
		return NULL;
	}
	return clazz;
}

const char* getCString(uintptr_t jni_env, uintptr_t ctx, jstring str) {
	JNIEnv *env = (JNIEnv*)jni_env;

	const char *chars = (*env)->GetStringUTFChars(env, str, NULL);

	const char *copy = strdup(chars);
	(*env)->ReleaseStringUTFChars(env, str, chars);
	return copy;
}

const char *androidName(uintptr_t java_vm, uintptr_t jni_env, uintptr_t ctx) {
	JNIEnv *env = (JNIEnv*)jni_env;

    // look up model from Build class
    jclass buildClass = find_class(env, "android/os/Build");
    jfieldID modelFieldID = (*env)->GetStaticFieldID(env, buildClass, "MODEL", "Ljava/lang/String;");
    jstring model = (*env)->GetStaticObjectField(env, buildClass, modelFieldID);

    // convert to a C string
	return getCString(jni_env, ctx, model);
}
