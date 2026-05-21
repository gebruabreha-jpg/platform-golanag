import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'core/routing/app_router.dart';
import 'core/theme/app_theme.dart';

void main() {
  runApp(const NiloCommerceApp());
}

class NiloCommerceApp extends StatelessWidget {
  const NiloCommerceApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
      providers: [
        // Add your BLoC providers here
      ],
      child: MaterialApp(
        title: 'Nilo Commerce',
        debugShowCheckedModeBanner: false,
        theme: AppTheme.lightTheme,
        darkTheme: AppTheme.darkTheme,
        onGenerateRoute: AppRouter.generateRoute,
        initialRoute: AppRouter.splash,
      ),
    );
  }
}
