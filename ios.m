//go:build ios

#import <UIKit/UIKit.h>

char *iosName() {
    NSString *model = [[UIDevice currentDevice] model];
    NSString *system = [[UIDevice currentDevice] systemName];
    NSString *version = [[UIDevice currentDevice] systemVersion];
    return [[NSString stringWithFormat:@"%@: %@ v%@", model, system, version] UTF8String];
}
