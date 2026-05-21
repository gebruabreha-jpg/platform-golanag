'use client';

import { useState } from 'react';
import Link from 'next/link';
import {
  Globe,
  Users,
  Shield,
  Zap,
  Home,
  Plane,
  CreditCard,
  Briefcase,
  Award,
  ArrowRight,
  CheckCircle,
  Star,
  MapPin,
  DollarSign,
  Clock,
  Heart,
  MessageCircle,
} from 'lucide-react';

export default function HomePage() {
  const [activeSection, setActiveSection] = useState('home');

  const services = [
    {
      id: 'shipping',
      title: 'Peer-to-Peer Shipping & Travel Delivery',
      description:
        'Connect with travelers heading to Ethiopia to send or receive items securely through unused luggage space.',
      icon: Plane,
      color: 'bg-blue-500',
      features: [
        'Find travelers matching your route & dates',
        'Secure escrow payments (pay on delivery)',
        'Track package status in real-time',
        'Verified traveler profiles & reviews',
      ],
      example: 'PiggyBee, Grabr, Worldee, Airlift Hub',
    },
    {
      id: 'marketplace',
      title: 'Buy & Sell Goods',
      description:
        'Ethiopia and diaspora marketplace for electronics, clothing, vehicles, and household items.',
      icon: CreditCard,
      color: 'bg-green-500',
      features: [
        'List items for sale with detailed descriptions',
        'AI-powered scam detection',
        'Integrated payment escrow',
        'Shipping coordination',
      ],
      example: 'Jiji Ethiopia, Qefira, Facebook Marketplace',
    },
    {
      id: 'housing',
      title: 'Community Housing & Shared Accommodation',
      description:
        'Find rooms, roommates, and apartments in Ethiopian cities and diaspora hubs worldwide.',
      icon: Home,
      color: 'bg-purple-500',
      features: [
        'Search by city, price, room type',
        'AI-powered roommate matching',
        'Verified landlord & tenant profiles',
        'Integrated lease management',
      ],
      example: 'Roommates.com, SpareRoom, Airbnb (Habesha groups)',
    },
    {
      id: 'scholarships',
      title: 'Scholarships & Study Abroad',
      description:
        'Discover legitimate scholarships, exchange programs, and immigration resources.',
      icon: Award,
      color: 'bg-yellow-500',
      features: [
        'Curated scholarship database',
        'Eligibility matching AI',
        'Application deadline tracking',
        'Immigration guidance resources',
      ],
      example:
        'Erasmus Mundus, DAAD, Chevening, Fulbright, Mastercard Foundation',
    },
    {
      id: 'currency',
      title: 'Peer-to-Peer Currency Exchange',
      description:
        'Connect with trusted diaspora members for direct USD/EUR/GBP to ETB exchanges at better rates.',
      icon: DollarSign,
      color: 'bg-emerald-500',
      features: [
        'Direct peer-to-peer rates (no remittance fees)',
        'Trust scoring system',
        'Secure meetup locations',
        'Transaction history & verification',
      ],
      example: 'Wise (official), plus community Telegram groups',
    },
    {
      id: 'jobs',
      title: 'Freelancing & Remote Work',
      description:
        'Find remote jobs and freelancing opportunities for Ethiopians and diaspora professionals.',
      icon: Briefcase,
      color: 'bg-red-500',
      features: [
        'Remote job listings worldwide',
        'Skill-based matching',
        'Freelancer portfolio profiles',
        'Secure contract payments',
      ],
      example: 'Upwork, Fiverr, Toptal, LinkedIn, Remote OK',
    },
    {
      id: 'networking',
      title: 'Business & Community Networking',
      description:
        'Connect with Ethiopian entrepreneurs, investors, and professionals globally.',
      icon: Users,
      color: 'bg-indigo-500',
      features: [
        'Find co-founders & business partners',
        'Discover investment opportunities',
        'Attend diaspora networking events',
        'Mentorship connections',
      ],
      example: 'Ethiopian Diaspora Trust Fund, Meetup, Clubhouse',
    },
  ];

  const trustFeatures = [
    {
      icon: Shield,
      title: 'Verified Profiles',
      description:
        'Multi-layer verification: ID, phone, social proof, transaction history',
    },
    {
      icon: Star,
      title: 'Trust Scoring System',
      description:
        'AI-powered trust score based on reviews, completed transactions, behavior patterns',
    },
    {
      icon: Clock,
      title: '24/7 Availability',
      description:
        'Platform available worldwide, community support in multiple time zones',
    },
    {
      icon: Heart,
      title: 'Community-Driven',
      description:
        'Built by Ethiopians for Ethiopians. Governance by community-elected council',
    },
    {
      icon: DollarSign,
      title: 'Low Cost',
      description:
        'No remittance fees. Minimal transaction fees (2-3%). Peer exchange rates better than banks',
    },
    {
      icon: Globe,
      title: 'Global Reach',
      description:
        'Connect from any country. Available in major diaspora hubs: US, Europe, Middle East',
    },
  ];

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Hero Section */}
      <header className="bg-white shadow-sm sticky top-0 z-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center py-4">
            <div className="flex items-center space-x-2">
              <div className="w-10 h-10 bg-gradient-to-br from-primary-500 to-purple-600 rounded-lg flex items-center justify-center">
                <Globe className="text-white" size={24} />
              </div>
              <span className="text-2xl font-bold text-gray-900">ConnectMe</span>
            </div>
            <nav className="hidden md:flex items-center space-x-6">
              {services.map((s) => (
                <a
                  key={s.id}
                  href={`#${s.id}`}
                  className="text-gray-600 hover:text-primary-600 font-medium"
                >
                  {s.title.split(' ')[0]}
                </a>
              ))}
              <Link
                href="/auth/login"
                className="text-gray-600 hover:text-primary-600 font-medium"
              >
                Sign In
              </Link>
              <Link
                href="/auth/register"
                className="bg-primary-600 text-white px-4 py-2 rounded-lg hover:bg-primary-700"
              >
                Join Free
              </Link>
            </nav>
          </div>
        </div>
      </header>

      {/* Hero */}
      <section className="bg-gradient-to-br from-primary-600 via-purple-600 to-pink-500 text-white py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h1 className="text-5xl md:text-6xl font-bold mb-6">
            Connect Ethiopians <br />
            <span className="text-yellow-300">Worldwide</span>
          </h1>
          <p className="text-xl md:text-2xl mb-10 max-w-3xl mx-auto opacity-90">
            The all-in-one platform for diaspora and locals: shipping, marketplace,
            housing, scholarships, jobs, currency exchange, and community networking.
          </p>
          <div className="flex flex-col sm:flex-row justify-center gap-4">
            <Link
              href="/auth/register"
              className="inline-flex items-center px-8 py-4 bg-white text-primary-600 text-lg font-bold rounded-xl hover:bg-gray-100 transition"
            >
              Get Started Free
              <ArrowRight className="ml-2" size={20} />
            </Link>
            <a
              href="#services"
              className="inline-flex items-center px-8 py-4 border-2 border-white text-white text-lg font-semibold rounded-xl hover:bg-white/10 transition"
            >
              Explore Services
            </a>
          </div>

          {/* Stats */}
          <div className="mt-16 grid grid-cols-2 md:grid-cols-4 gap-8">
            {[
              { value: '50K+', label: 'Active Users' },
              { value: '120K+', label: 'Successful Shipments' },
              { value: '$2M+', label: 'Transactions Processed' },
              { value: '98%', label: 'Satisfaction Rate' },
            ].map((stat) => (
              <div key={stat.label} className="text-center">
                <div className="text-3xl md:text-4xl font-bold">{stat.value}</div>
                <div className="text-sm opacity-75">{stat.label}</div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Services Section */}
      <section id="services" className="py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-4xl font-bold text-gray-900 mb-4">
              7 Essential Services in One Platform
            </h2>
            <p className="text-xl text-gray-600">
              Everything Ethiopians and diaspora need to connect, transact, and thrive
            </p>
          </div>

          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
            {services.map((service) => (
              <div
                key={service.id}
                id={service.id}
                className="card group hover:shadow-xl transition-all duration-300"
              >
                <div className={`w-14 h-14 ${service.color} rounded-xl flex items-center justify-center mb-4`}>
                  <service.icon className="text-white" size={28} />
                </div>
                <h3 className="text-xl font-bold mb-3">{service.title}</h3>
                <p className="text-gray-600 mb-4">{service.description}</p>
                <ul className="space-y-2 mb-4">
                  {service.features.map((feature, idx) => (
                    <li key={idx} className="flex items-start space-x-2 text-sm">
                      <CheckCircle className="text-green-500 flex-shrink-0 mt-0.5" size={16} />
                      <span>{feature}</span>
                    </li>
                  ))}
                </ul>
                <div className="pt-4 border-t border-gray-100">
                  <p className="text-sm text-gray-500 mb-2">Inspired by: {service.example}</p>
                  <Link
                    href={`/${service.id}`}
                    className="inline-flex items-center text-primary-600 font-semibold group-hover:underline"
                  >
                    Try it now <ArrowRight className="ml-1" size={16} />
                  </Link>
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Trust & Why Section */}
      <section className="py-20 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-4xl font-bold text-gray-900 mb-4">
              Why Trust ConnectMe?
            </h2>
            <p className="text-xl text-gray-600">
              Built for trust, transparency, and community
            </p>
          </div>

          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
            {trustFeatures.map((feature, idx) => (
              <div key={idx} className="flex items-start space-x-4 p-6 rounded-lg bg-gray-50">
                <div className="flex-shrink-0 w-12 h-12 bg-primary-100 rounded-full flex items-center justify-center">
                  <feature.icon className="text-primary-600" size={24} />
                </div>
                <div>
                  <h3 className="font-bold text-gray-900 mb-1">{feature.title}</h3>
                  <p className="text-gray-600 text-sm">{feature.description}</p>
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Cost Section */}
      <section className="py-20 bg-green-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-4xl font-bold text-gray-900 mb-4">Low Cost, High Value</h2>
            <p className="text-xl text-gray-600">Keep more money in your pocket</p>
          </div>

          <div className="grid md:grid-cols-3 gap-8">
            <div className="bg-white rounded-2xl shadow-lg p-8 text-center">
              <DollarSign className="w-12 h-12 text-green-500 mx-auto mb-4" />
              <h3 className="text-2xl font-bold mb-2">Platform Fee</h3>
              <p className="text-4xl font-bold text-primary-600 mb-2">2%</p>
              <p className="text-gray-600">Per transaction (buyer or seller)</p>
              <p className="text-sm text-gray-500 mt-4">
                Compared to 5-10% typical remittance fees
              </p>
            </div>

            <div className="bg-white rounded-2xl shadow-lg p-8 text-center">
              <Globe className="w-12 h-12 text-blue-500 mx-auto mb-4" />
              <h3 className="text-2xl font-bold mb-2">FX Spread</h3>
              <p className="text-4xl font-bold text-primary-600 mb-2">0.5%</p>
              <p className="text-gray-600">Mid-market rate markup</p>
              <p className="text-sm text-gray-500 mt-4">
                Better than banks (2-4%) and money transfer services
              </p>
            </div>

            <div className="bg-white rounded-2xl shadow-lg p-8 text-center">
              <Heart className="w-12 h-12 text-red-500 mx-auto mb-4" />
              <h3 className="text-2xl font-bold mb-2">Community Benefit</h3>
              <p className="text-2xl font-bold text-primary-600 mb-2">Free</p>
              <p className="text-gray-600">Groups, events, networking</p>
              <p className="text-sm text-gray-500 mt-4">
                Community features always free for members
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* Requirements Section */}
      <section className="py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-4xl font-bold text-gray-900 mb-4">Getting Started is Easy</h2>
            <p className="text-xl text-gray-600">Simple requirements, powerful platform</p>
          </div>

          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-8">
            {[
              {
                step: '01',
                title: 'Create Account',
                desc: 'Sign up with email or phone. Verify with ID for full access.',
              },
              {
                step: '02',
                title: 'Complete Profile',
                desc: 'Add location, interests, skills, and preferences for better matching.',
              },
              {
                step: '03',
                title: 'Get Verified',
                desc: 'Submit ID and proof of address. Higher trust = more features.',
              },
              {
                step: '04',
                title: 'Start Connecting',
                desc: 'Browse communities, list items, find housing, get jobs.',
              },
            ].map((step) => (
              <div key={step.step} className="relative">
                <div className="text-6xl font-bold text-gray-100 absolute top-0 left-0">
                  {step.step}
                </div>
                <div className="relative z-10">
                  <h3 className="text-lg font-bold mb-2">{step.title}</h3>
                  <p className="text-gray-600">{step.desc}</p>
                </div>
              </div>
            ))}
          </div>

          <div className="mt-12 bg-yellow-50 border-l-4 border-yellow-400 p-6 rounded-lg">
            <div className="flex items-start">
              <Shield className="text-yellow-500 flex-shrink-0 mt-1" size={24} />
              <div className="ml-4">
                <h3 className="font-bold text-gray-900">Important Legal Notice</h3>
                <p className="mt-2">
                  ConnectMe is a community platform. All users must comply with local and international laws.
                  Fraud, scams, or illegal activities will result in immediate account termination and legal
                  prosecution. Never share sensitive financial information outside secure payment systems.
                </p>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-20 bg-gradient-to-br from-primary-600 to-purple-700 text-white">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h2 className="text-4xl md:text-5xl font-bold mb-6">
            Join 50,000+ Ethiopians Worldwide
          </h2>
          <p className="text-xl mb-10 opacity-90">
            Connect, transact, and build relationships with the global Ethiopian community.
            Safe, trusted, and built for you.
          </p>
          <div className="flex flex-col sm:flex-row justify-center gap-4">
            <Link
              href="/auth/register"
              className="inline-flex items-center px-8 py-4 bg-white text-primary-600 text-lg font-bold rounded-xl hover:bg-gray-100 transition"
            >
              Create Free Account
              <ArrowRight className="ml-2" size={20} />
            </Link>
            <Link
              href="/about"
              className="inline-flex items-center px-8 py-4 border-2 border-white text-white text-lg font-semibold rounded-xl hover:bg-white/10 transition"
            >
              Learn More
            </Link>
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="bg-gray-900 text-gray-300 py-12">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 grid grid-cols-1 md:grid-cols-4 gap-8">
          <div>
            <div className="flex items-center space-x-2 mb-4">
              <div className="w-10 h-10 bg-white rounded-lg flex items-center justify-center">
                <span className="text-primary-600 font-bold text-xl">C</span>
              </div>
              <span className="text-2xl font-bold text-white">ConnectMe</span>
            </div>
            <p className="text-sm">
              Empowering Ethiopian connectivity worldwide through trusted community services.
            </p>
          </div>
          <div>
            <h4 className="font-semibold text-white mb-4">Services</h4>
            <ul className="space-y-2 text-sm">
              <li><a href="#" className="hover:text-white">Shipping</a></li>
              <li><a href="#" className="hover:text-white">Marketplace</a></li>
              <li><a href="#" className="hover:text-white">Housing</a></li>
              <li><a href="#" className="hover:text-white">Scholarships</a></li>
            </ul>
          </div>
          <div>
            <h4 className="font-semibold text-white mb-4">Support</h4>
            <ul className="space-y-2 text-sm">
              <li><a href="#" className="hover:text-white">Help Center</a></li>
              <li><a href="#" className="hover:text-white">Safety Tips</a></li>
              <li><a href="#" className="hover:text-white">Contact Us</a></li>
              <li><a href="#" className="hover:text-white">Report Issue</a></li>
            </ul>
          </div>
          <div>
            <h4 className="font-semibold text-white mb-4">Legal</h4>
            <ul className="space-y-2 text-sm">
              <li><a href="#" className="hover:text-white">Terms of Service</a></li>
              <li><a href="#" className="hover:text-white">Privacy Policy</a></li>
              <li><a href="#" className="hover:text-white">Cookie Policy</a></li>
            </ul>
          </div>
        </div>
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 mt-8 pt-8 border-t border-gray-800 text-center text-sm">
          © 2026 ConnectMe. All rights reserved.
        </div>
      </footer>
    </div>
  );
}
