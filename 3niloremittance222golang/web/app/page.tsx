import Link from 'next/link';
import { ArrowRight, Globe, Shield, Zap, Users } from 'lucide-react';

export default function HomePage() {
  return (
    <div className="min-h-screen bg-white">
      {/* Header */}
      <header className="border-b border-gray-200">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center py-6">
            <div className="flex items-center space-x-2">
              <div className="w-10 h-10 bg-primary-600 rounded-lg flex items-center justify-center">
                <span className="text-white font-bold text-xl">N</span>
              </div>
              <span className="text-2xl font-bold text-gray-900">Nilo Commerce</span>
            </div>
            <div className="flex items-center space-x-4">
              <Link href="/login" className="text-gray-600 hover:text-primary-600">
                Sign In
              </Link>
              <Link
                href="/register"
                className="bg-primary-600 text-white px-6 py-2 rounded-lg hover:bg-primary-700 transition"
              >
                Get Started
              </Link>
            </div>
          </div>
        </div>
      </header>

      {/* Hero Section */}
      <main>
        <section className="py-20 bg-gradient-to-b from-primary-50 to-white">
          <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
            <h1 className="text-5xl font-bold text-gray-900 mb-6">
              Direct Bill Payments for Ethiopia
            </h1>
            <p className="text-xl text-gray-600 mb-10 max-w-3xl mx-auto">
              The secure, legal way for the Ethiopian diaspora to pay for school tuition, healthcare,
              and grocery vouchers directly to merchants in Ethiopia. No remittance. No cash transfers.
            </p>
            <div className="flex justify-center gap-4">
              <Link
                href="/register"
                className="inline-flex items-center px-8 py-4 bg-primary-600 text-white text-lg font-semibold rounded-xl hover:bg-primary-700 transition"
              >
                Start Paying Directly
                <ArrowRight className="ml-2" size={20} />
              </Link>
              <Link
                href="/merchants"
                className="inline-flex items-center px-8 py-4 border-2 border-gray-300 text-gray-700 text-lg font-semibold rounded-xl hover:border-primary-600 hover:text-primary-600 transition"
              >
                Register as Merchant
              </Link>
            </div>
          </div>
        </section>

        {/* How It Works */}
        <section className="py-20">
          <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <h2 className="text-3xl font-bold text-center mb-12">How It Works</h2>
            <div className="grid md:grid-cols-3 gap-8">
              <div className="text-center p-6">
                <div className="w-16 h-16 bg-primary-100 rounded-full flex items-center justify-center mx-auto mb-4">
                  <Globe className="text-primary-600" size={32} />
                </div>
                <h3 className="text-xl font-semibold mb-2">1. Diaspora Pays Online</h3>
                <p className="text-gray-600">
                  Family abroad uses our platform to pay for school fees, hospital bills, or grocery vouchers with Stripe.
                </p>
              </div>
              <div className="text-center p-6">
                <div className="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-4">
                  <Zap className="text-green-600" size={32} />
                </div>
                <h3 className="text-xl font-semibold mb-2">2. We Process & Convert</h3>
                <p className="text-gray-600">
                  Funds arrive at our Wise Business account. We convert USD to ETB and handle all compliance.
                </p>
              </div>
              <div className="text-center p-6">
                <div className="w-16 h-16 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-4">
                  <Shield className="text-blue-600" size={32} />
                </div>
                <h3 className="text-xl font-semibold mb-2">3. Merchant Gets Paid</h3>
                <p className="text-gray-600">
                  The school, hospital, or supermarket receives a local Ethiopian bank transfer directly.
                </p>
              </div>
            </div>
          </div>
        </section>

        {/* Benefits Section */}
        <section className="py-20 bg-gray-50">
          <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <h2 className="text-3xl font-bold text-center mb-12">Why Nilo Commerce?</h2>
            <div className="grid md:grid-cols-2 gap-8">
              <div className="flex items-start space-x-4">
                <Shield className="text-green-500 flex-shrink-0" size={24} />
                <div>
                  <h3 className="font-semibold text-lg">100% Legal & Compliant</h3>
                  <p className="text-gray-600">
                    Operates as a bill payment platform under the "agent of the payee" exception. Fully compliant with FinCEN regulations.
                  </p>
                </div>
              </div>
              <div className="flex items-start space-x-4">
                <Users className="text-blue-500 flex-shrink-0" size={24} />
                <div>
                  <h3 className="font-semibold text-lg">Zero Remittance Risk</h3>
                  <p className="text-gray-600">
                    Money flows directly from payer to merchant via closed-loop system. No holding or transmitting customer funds.
                  </p>
                </div>
              </div>
              <div className="flex items-start space-x-4">
                <Zap className="text-yellow-500 flex-shrink-0" size={24} />
                <div>
                  <h3 className="font-semibold text-lg">Fast Settlements</h3>
                  <p className="text-gray-600">
                    Merchants receive local ETB bank transfers within 24-48 hours of payment.
                  </p>
                </div>
              </div>
              <div className="flex items-start space-x-4">
                <Globe className="text-purple-500 flex-shrink-0" size={24} />
                <div>
                  <h3 className="font-semibold text-lg">Wise & Stripe Powered</h3>
                  <p className="text-gray-600">
                    Built on trusted financial infrastructure. Secure, reliable, and globally recognized.
                  </p>
                </div>
              </div>
            </div>
          </div>
        </section>

        {/* Call to Action */}
        <section className="py-20 bg-primary-600 text-white">
          <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
            <h2 className="text-3xl font-bold mb-4">
              Ready to Simplify Cross-Border Payments?
            </h2>
            <p className="text-xl mb-8 text-primary-100">
              Join hundreds of Ethiopian merchants and diaspora families already using Nilo Commerce.
            </p>
            <Link
              href="/register"
              className="inline-flex items-center px-8 py-4 bg-white text-primary-600 text-lg font-semibold rounded-xl hover:bg-gray-100 transition"
            >
              Get Started Today
              <ArrowRight className="ml-2" size={20} />
            </Link>
          </div>
        </section>
      </main>

      {/* Footer */}
      <footer className="bg-gray-900 text-gray-300 py-12">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 grid grid-cols-1 md:grid-cols-4 gap-8">
          <div>
            <div className="flex items-center space-x-2 mb-4">
              <div className="w-10 h-10 bg-white rounded-lg flex items-center justify-center">
                <span className="text-primary-600 font-bold text-xl">N</span>
              </div>
              <span className="text-2xl font-bold text-white">Nilo Commerce</span>
            </div>
            <p className="text-sm">
              Closed-loop B2B2C bill payment platform connecting the Ethiopian diaspora with local merchants.
            </p>
          </div>
          <div>
            <h4 className="font-semibold text-white mb-4">For Diaspora</h4>
            <ul className="space-y-2 text-sm">
              <li><Link href="/how-it-works" className="hover:text-white">How It Works</Link></li>
              <li><Link href="/merchants" className="hover:text-white">Find Merchants</Link></li>
              <li><Link href="/pricing" className="hover:text-white">Pricing</Link></li>
              <li><Link href="/help" className="hover:text-white">Help Center</Link></li>
            </ul>
          </div>
          <div>
            <h4 className="font-semibold text-white mb-4">For Merchants</h4>
            <ul className="space-y-2 text-sm">
              <li><Link href="/merchant/register" className="hover:text-white">Register</Link></li>
              <li><Link href="/merchant/benefits" className="hover:text-white">Benefits</Link></li>
              <li><Link href="/merchant/pricing" className="hover:text-white">Merchant Pricing</Link></li>
              <li><Link href="/merchant/faq" className="hover:text-white">Merchant FAQ</Link></li>
            </ul>
          </div>
          <div>
            <h4 className="font-semibold text-white mb-4">Company</h4>
            <ul className="space-y-2 text-sm">
              <li><Link href="/about" className="hover:text-white">About</Link></li>
              <li><Link href="/contact" className="hover:text-white">Contact</Link></li>
              <li><Link href="/legal" className="hover:text-white">Legal</Link></li>
              <li><Link href="/privacy" className="hover:text-white">Privacy</Link></li>
            </ul>
          </div>
        </div>
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 mt-8 pt-8 border-t border-gray-800 text-center text-sm">
          © 2026 Nilo Commerce. All rights reserved.
        </div>
      </footer>
    </div>
  );
}
