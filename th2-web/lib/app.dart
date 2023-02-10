import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:hive/hive.dart';
import 'package:th2_web/routes/app_route.dart';
import 'package:th2_web/utils/custom_scroll_behavior.dart';
import 'package:th2_web/utils/extensions.dart';
import 'package:th2_web/view_models/app_bloc/app_bloc.dart';
import 'package:th2_web/view_models/bloc_wrapper.dart';
import 'package:th2_web/views/common_widgets/footer.dart';

import 'generated/l10n.dart';

class ReApp extends StatefulWidget {
  const ReApp({Key? key}) : super(key: key);

  @override
  State<ReApp> createState() => _ReAppState();
}

class _ReAppState extends State<ReApp> {
  @override
  void dispose() {
    Hive.close();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return BlocWrapper(
      child: Builder(builder: (context) {
        return GestureDetector(
          onTap: () {
            FocusManager.instance.primaryFocus?.unfocus();
          },
          child: ScreenUtilInit(
            //Put Figma size here
            designSize: const Size(1920, 1080),
            builder: (context, child) {
              return child!;
            },
            child: MaterialApp.router(
              scrollBehavior: MyCustomScrollBehavior(),
              localizationsDelegates: const [
                S.delegate,
                GlobalMaterialLocalizations.delegate,
                GlobalWidgetsLocalizations.delegate,
                GlobalCupertinoLocalizations.delegate,
              ],
              supportedLocales: S.delegate.supportedLocales,
              locale: context.watch<AppBloc>().state.locale,
              debugShowCheckedModeBanner: false,
              routerConfig: router,
              builder: (context, widget) {
                if (widget == null) {
                  debugPrint("Material builder: widget is null");
                  return const SizedBox();
                }
                return Stack(
                  children: [
                    Positioned.fill(child: widget),
                    Positioned.fill(
                      child: Align(
                        alignment: Alignment.bottomCenter,
                        child: const Footer()
                            .withPadding(const EdgeInsets.only(bottom: 16)),
                      ),
                    ),
                  ],
                );
              },
            ),
          ),
        );
      }),
    );
  }
}
