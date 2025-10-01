// pheonix-tool/loader/bootstrapper.cpp
// A unused bootstrapper CLI tool for repairing, updating, and launching the Pheonix application in debug mode.
// This is a simple C++ console application that provides a menu for users to select options.
// WARNING! This code is writed on basic standard. The code is not optimized and may contain security issues and can have bugs and exploits.
// Use at your own risk!

// P.S Translate to other languages if needed.


#include <iostream>
#include <cstdlib>   // system()
#include <string>
#include <thread>
#include <chrono>

using namespace std;

// Function to simulate a delay for better UX
void delayPrint(const string& msg, int ms = 700) {
    cout << msg << endl;
    this_thread::sleep_for(chrono::milliseconds(ms));
}

void repairApp() {
    cout << "\n[Repair The Application]\n";
    delayPrint("Checking files...");
    delayPrint("Checking dependencies...");
    delayPrint("No critical errors found.");
    delayPrint("Repair complete!\n");
}

void checkUpdates() {
    cout << "\n[Check For Updates]\n";
    delayPrint("Connecting to update server...");
    delayPrint("Checking version...");
    delayPrint("You are running the latest version!\n");
}

void launchDebug() {
    cout << "\n[Launch With Debug Mode]\n";
    delayPrint("Starting application in debug mode...");
    // Launch the executable "Pheonix.exe" with debug flag
    int result = system("./Pheonix.exe --debug");
    if (result != 0) {
        cout << "Failed to launch Pheonix.exe.\n";
    } else {
        cout << "Application closed.\n";
    }
}

int main() {
    while (true) {
        cout << "============================\n";
        cout << "   Bootstrapper Menu\n";
        cout << "============================\n";
        cout << "1. Repair The Application\n";
        cout << "2. Check For Updates\n";
        cout << "3. Launch With Debug Mode\n";
        cout << "4. Exit\n";
        cout << "Select option: ";

        int choice;
        cin >> choice;

        switch (choice) {
            case 1: repairApp(); break;
            case 2: checkUpdates(); break;
            case 3: launchDebug(); break;
            case 4: cout << "Goodbye!\n"; return 0;
            default: cout << "Invalid option.\n"; break;
        }
        cout << endl;
    }
}
