import pandas as pd
import matplotlib.pyplot as plt


# Function to classify the correlation
def classify_correlation(correlation):
    if correlation == 1.0:
        return "Perfect Positive Correlation"
    elif 0.7 <= correlation < 1.0:
        return "Strong Positive Correlation"
    elif 0.3 <= correlation < 0.7:
        return "Moderate Positive Correlation"
    elif 0.1 <= correlation < 0.3:
        return "Weak Positive Correlation"
    elif correlation == 0.0:
        return "No Correlation"
    elif -0.1 <= correlation < 0.0:
        return "Weak Negative Correlation"
    elif -0.3 <= correlation < -0.1:
        return "Moderate Negative Correlation"
    elif -0.7 <= correlation < -0.3:
        return "Strong Negative Correlation"
    elif correlation == -1.0:
        return "Perfect Negative Correlation"
    else:
        return "Undefined Correlation"



# Load the data
data = pd.read_csv('weather_data.csv')

correlation = data['Temperature'].corr(data['AirPressure'])

classification = classify_correlation(correlation)
print("Correlation Classification: ", classification)
print("correlation: ", round(correlation, 2))






# Plotting
plt.figure(figsize=(10, 6))
plt.scatter(data['AirPressure'], data['Temperature'], color='blue', alpha=0.5)
plt.title('Temperature vs Air Pressure')
plt.xlabel('Air Pressure (hPa)')
plt.ylabel('Temperature (K)')
plt.grid(True)
#plt.show()
plt.savefig('analysis.png')


