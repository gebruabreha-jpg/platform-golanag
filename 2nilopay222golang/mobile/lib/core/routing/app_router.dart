import 'package:flutter/material.dart';
import '../presentation/screens/splash/splash_screen.dart';
import '../presentation/screens/auth/login_screen.dart';
import '../presentation/screens/auth/register_screen.dart';
import '../presentation/screens/home/home_screen.dart';
import '../presentation/screens/merchants/merchant_list_screen.dart';
import '../presentation/screens/transactions/transaction_screen.dart';
import '../presentation/screens/transactions/transaction_confirmation_screen.dart';
import '../presentation/screens/transactions/transaction_history_screen.dart';

class AppRouter {
  static const String splash = '/';
  static const String login = '/login';
  static const String register = '/register';
  static const String home = '/home';
  static const String merchants = '/merchants';
  static const String makePayment = '/make-payment';
  static const String confirmPayment = '/confirm-payment';
  static const String history = '/history';

  static Route<dynamic> generateRoute(RouteSettings settings) {
    switch (settings.name) {
      case splash:
        return MaterialPageRoute(builder: (_) => const SplashScreen());
      case login:
        return MaterialPageRoute(builder: (_) => const LoginScreen());
      case register:
        return MaterialPageRoute(builder: (_) => const RegisterScreen());
      case home:
        return MaterialPageRoute(builder: (_) => const HomeScreen());
      case merchants:
        return MaterialPageRoute(builder: (_) => const MerchantListScreen());
      case makePayment:
        final args = settings.arguments as Map<String, dynamic>?;
        return MaterialPageRoute(
          builder: (_) => TransactionScreen(merchant: args?['merchant']),
        );
      case confirmPayment:
        final args = settings.arguments as Map<String, dynamic>;
        return MaterialPageRoute(
          builder: (_) => TransactionConfirmationScreen(
            transaction: args['transaction'],
          ),
        );
      case history:
        return MaterialPageRoute(builder: (_) => const TransactionHistoryScreen());
      default:
        return MaterialPageRoute(
          builder: (_) => Scaffold(
            body: Center(
              child: Text('No route defined for ${settings.name}'),
            ),
          ),
        );
    }
  }
}
