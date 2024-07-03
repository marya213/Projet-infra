import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonArray;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;

import javax.swing.*;
import javax.swing.text.*;
import java.awt.*;
import java.awt.datatransfer.*;
import java.awt.dnd.*;
import java.awt.event.*;
import java.io.*;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;

public class TextFieldTest implements ActionListener {
    private static final String DEFAULT_FILE_PATH = "resources/task1.json";
    private static final Gson gson = new GsonBuilder().setPrettyPrinting().create();
    private List<JsonObject> tasks; 
    private String currentFilePath;
    private JTextField text1;
    private JTextPane text2;
    private JButton btnAdd;
    private JButton btnOpen;
    private JButton btnDelete;
    private JButton btnEdit;
    private JTextField txtDate;
    private JTextField txtTime;
    private JComboBox<String> fileListComboBox;
    private static TextFieldTest instance;
// Get instance
    public static TextFieldTest getInstance() {
        if (instance == null) {
            instance = new TextFieldTest();
        }
        return instance;
    }
// Constructor
    private TextFieldTest() {
        tasks = new ArrayList<>();
        currentFilePath = DEFAULT_FILE_PATH;
        createUI();
        loadTasksFromFile(currentFilePath);
    }
// Create the UI
    private void createUI() {
        JFrame f = new JFrame("Welcome Task Manager");
        f.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        f.setLayout(new BorderLayout());

        text1 = new JTextField();
        text1.setPreferredSize(new Dimension(300, 30));
// Add the text field
        JPanel buttonPanel = new JPanel(new FlowLayout(FlowLayout.CENTER));
        btnAdd = createButton("Add task");
        btnOpen = createButton("Open task");
        btnEdit = createButton("Edit task");
        btnDelete = createButton("Delete task");
        buttonPanel.add(btnAdd);
        buttonPanel.add(btnOpen);
        buttonPanel.add(btnEdit);
        buttonPanel.add(btnDelete);
// Add the file list combo box
        JPanel filterPanel = new JPanel(new FlowLayout(FlowLayout.CENTER));
        JComboBox<String> categoryFilter = new JComboBox<>(new String[]{"All", "Home", "Work", "Personal"});
        JComboBox<String> priorityFilter = new JComboBox<>(new String[]{"All", "High", "Medium", "Low"});
        filterPanel.add(new JLabel("Filter by Category:"));
        filterPanel.add(categoryFilter);
        filterPanel.add(new JLabel("Filter by Priority:"));
        filterPanel.add(priorityFilter);
// Add the text pane
        text2 = createTextPane();
        text2.setBackground(new Color(240, 240, 240));
        JScrollPane scrollPane = new JScrollPane(text2);
        scrollPane.setPreferredSize(new Dimension(400, 300));
        f.add(scrollPane, BorderLayout.CENTER);
        f.add(buttonPanel, BorderLayout.SOUTH);
        f.add(filterPanel, BorderLayout.NORTH);
// Add the file list combo box
        fileListComboBox = new JComboBox<>(new String[]{"task1.json"});
        fileListComboBox.addItemListener(new ItemListener() {
            public void itemStateChanged(ItemEvent e) {
                if (e.getStateChange() == ItemEvent.SELECTED) {
                    String selectedFile = (String) e.getItem();
                    currentFilePath = "resources/" + selectedFile;
                    loadTasksFromFile(currentFilePath);
                }
            }
        });
//  Add the file list combo box
        categoryFilter.addActionListener(e -> refreshTaskDisplay(categoryFilter.getSelectedItem().toString(), priorityFilter.getSelectedItem().toString()));
        priorityFilter.addActionListener(e -> refreshTaskDisplay(categoryFilter.getSelectedItem().toString(), priorityFilter.getSelectedItem().toString()));

        f.pack();
        f.setLocationRelativeTo(null);
        f.setVisible(true);
        text2.setDropTarget(new DropTarget(text2, new TaskDropTargetListener()));
    }
 // Create the text pane
    private JTextPane createTextPane() {
        JTextPane textPane = new JTextPane();
        textPane.setEditable(false);
        StyledDocument doc = textPane.getStyledDocument();
        Style style = doc.addStyle("Style", null);
        StyleConstants.setForeground(style, Color.BLACK);
        StyleConstants.setFontSize(style, 14);
        textPane.setParagraphAttributes(style, true);
        return textPane;
    }
// Create the button
    private JButton createButton(String label) {
        JButton button = new JButton(label);
        button.addActionListener(this);
        return button;
    }
// Action performed
    public void actionPerformed(ActionEvent e) {
        if (e.getSource() == btnAdd) {
            openAddTaskDialog();
        } else if (e.getSource() == btnOpen) {
            openTask();
        } else if (e.getSource() == btnEdit) {
            openEditDialog();
        } else if (e.getSource() == btnDelete) {
            openDeleteDialog();
        }
    }
// Open the add task dialog
    private void openAddTaskDialog() {
        AddTaskDialog dialog = new AddTaskDialog((JFrame) SwingUtilities.getWindowAncestor(text1), this);
        dialog.setVisible(true);
    }
// Add task from dialog
    public void addTaskFromDialog(String name, String description, String priority, String category, String dateTime) {
    JsonObject taskObject = createTaskJson(name, description, priority, category, dateTime);
    tasks.add(taskObject);
    saveTasksToFile();
    refreshTaskDisplay("All", "All");
    text1.setText("");
}

    // Open the task
    private void openTask() {
        String taskName = getUserInput("Enter the name of the task:");
        if (taskName != null && !taskName.isEmpty()) {
            displayTaskDetails(taskName);
        }
    }
    // Display the task details
    private void displayTaskDetails(String taskName) {
        JsonObject taskObject = findTaskByTaskTitle(taskName);
        if (taskObject != null) {
            String taskDescription = taskObject.get("description").getAsString();
            showMessage("Description of the task : " + taskDescription, "Details of the task");
        } else {
            showMessage("Task not found.", taskName);
        }
    }
    
    // Find the task by task title
    private JsonObject findTaskByTaskTitle(String taskName) {
        for (JsonObject taskObject : tasks) {
            String taskTitle = taskObject.get("tasktitle").getAsString();
            if (taskTitle.equals(taskName)) {
                return taskObject;
            }
        }
        return null;
    }
    
    // Open the edit dialog
    private void openEditDialog() {
        String description = getUserInput("Enter the description of the task to edit :");
        if (description != null && !description.isEmpty()) {
            displayEditDialog(description);
        }
    }
    // Display the edit dialog
    private void displayEditDialog(String taskName) {
        JsonObject taskObject = findTaskByTaskTitle(taskName);
        if (taskObject != null) {
            String description = taskObject.get("description").getAsString();
            String priority = taskObject.get("priority").getAsString();
            String category = taskObject.get("category").getAsString();
            TaskEditDialog editDialog = new TaskEditDialog((JFrame) SwingUtilities.getWindowAncestor(text1), taskName, description, priority, category, this);
            editDialog.setVisible(true);
        } else {
            showMessage("Task not found.", taskName);
        }
    }
    // Open the delete dialog
    private void openDeleteDialog() {
        String description = getUserInput("Enter the description of the task to delete :");
        if (description != null && !description.isEmpty()) {
            int option = showConfirmation("Do you want to delete this task?", "Confirm Deletion");
            if (option == JOptionPane.YES_OPTION) {
                deleteTask(description);
            }
        }
    }
    // Delete the task
    private void deleteTask(String taskTitle) {
        for (int i = 0; i < tasks.size(); i++) {
            JsonObject taskObject = tasks.get(i);
            if (taskObject.get("tasktitle").getAsString().equals(taskTitle)) {
                tasks.remove(i);
                saveTasksToFile();
                refreshTaskDisplay("All", "All");
                break;
            }
        }
    }
  // Load tasks from file  
    private void loadTasksFromFile(String filePath) {
        tasks.clear();
        try (Reader reader = new FileReader(filePath)) {
            JsonObject jsonObject = gson.fromJson(reader, JsonObject.class);
            if (jsonObject.has("tasks")) {  
                JsonArray taskArray = jsonObject.getAsJsonArray("tasks");
                for (JsonElement taskElement : taskArray) {
                    tasks.add(taskElement.getAsJsonObject());
                }
            }
            refreshTaskDisplay("All", "All");
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
    // Get user input
    private String getUserInput(String message) {
        return JOptionPane.showInputDialog(null, message);
    }
    // Show message
    private void showMessage(String message, String title) {
        JOptionPane.showMessageDialog(null, message, title, JOptionPane.INFORMATION_MESSAGE);
    }
    // Show confirmation
    private int showConfirmation(String message, String title) {
        return JOptionPane.showConfirmDialog(null, message, title, JOptionPane.YES_NO_OPTION);
    }
    
// Create task json
private JsonObject createTaskJson(String name, String description, String priority, String category, String dateTime) {
    JsonObject taskObject = new JsonObject();
    taskObject.addProperty("tasktitle", name);
    taskObject.addProperty("description", description);
    taskObject.addProperty("priority", priority);
    taskObject.addProperty("category", category);
    taskObject.addProperty("addedOn", dateTime);
    return taskObject;
}
// Get color for priority
private Color getColorForPriority(String priority) {
    switch (priority) {
        case "High":
            return Color.RED;
        case "Medium":
            return Color.ORANGE;
        case "Low":
            return Color.GREEN;
        default:
            return Color.BLACK;
    }
}
// Save tasks to file
private void saveTasksToFile() {
    try (Writer writer = new FileWriter(currentFilePath)) {
        JsonObject jsonObject = new JsonObject();
        JsonArray jsonArray = new JsonArray();
        tasks.forEach(jsonArray::add);
        jsonObject.add("tasks", jsonArray);
        gson.toJson(jsonObject, writer);
    } catch (IOException e) {
        e.printStackTrace();
    }
}
// Append to text pane
private void appendToTextPane(String text, Color color) {
    StyledDocument doc = text2.getStyledDocument();
    SimpleAttributeSet keyWord = new SimpleAttributeSet();
    StyleConstants.setForeground(keyWord, color);
    try {
        doc.insertString(doc.getLength(), text, keyWord);
    } catch (BadLocationException e) {
        e.printStackTrace();
    }
}
// Refresh task display
private void refreshTaskDisplay(String categoryFilter, String priorityFilter) {
    text2.setText(""); // Clear existing text
    for (JsonObject taskObject : tasks) {
        String taskCategory = taskObject.get("category").getAsString();
        String taskPriority = taskObject.get("priority").getAsString();

        if ((categoryFilter.equals("All") || taskCategory.equals(categoryFilter)) &&
                (priorityFilter.equals("All") || taskPriority.equals(priorityFilter))) {
            String taskDisplay = "Name : " + taskObject.get("tasktitle").getAsString() + "\n";
            appendToTextPane(taskDisplay, getColorForPriority(taskPriority));
        }
    }
}
// Task drop target listener
private class TaskDropTargetListener implements DropTargetListener {
    @Override
    public void dragEnter(DropTargetDragEvent dtde) {
        dtde.acceptDrag(DnDConstants.ACTION_COPY_OR_MOVE);
    }

    @Override
    public void dragOver(DropTargetDragEvent dtde) {
    }

    @Override
    public void dropActionChanged(DropTargetDragEvent dtde) {
    }

    @Override
    public void dragExit(DropTargetEvent dte) {
    }

    @Override
    public void drop(DropTargetDropEvent dtde) {
        Transferable transferable = dtde.getTransferable();
        if (transferable.isDataFlavorSupported(DataFlavor.stringFlavor)) {
            dtde.acceptDrop(DnDConstants.ACTION_COPY_OR_MOVE);
            try {
                String droppedText = (String) transferable.getTransferData(DataFlavor.stringFlavor);
                appendToTextPane("\nText added : " + droppedText, Color.BLUE);
                dtde.dropComplete(true);
            } catch (UnsupportedFlavorException | IOException e) {
                dtde.dropComplete(false);
            }
        } else {
            dtde.rejectDrop();
        }
    }
}
// Add task dialog
private class AddTaskDialog extends JDialog {
    private JTextField txtName;
    private JTextField txtDescription;
    private JComboBox<String> cmbPriority;
    private JComboBox<String> cmbCategory;
    private JButton btnAddCategory;
    private JTextField txtNewCategory;
    private TextFieldTest parent;
// Constructor
public AddTaskDialog(JFrame parent, TextFieldTest parentInstance) {
    super(parent, "Add Task", true);
    setSize(300, 250); // Adjusted size to accommodate new fields
    setLocationRelativeTo(parent);
    setLayout(new GridLayout(7, 2)); // Adjusted layout to accommodate new fields
    this.parent = parentInstance;
    JLabel lblName = new JLabel("Name:");
    txtName = new JTextField();
    JLabel lblDescription = new JLabel("Description:");
    txtDescription = new JTextField();
    JLabel lblPriority = new JLabel("Priority:");
    cmbPriority = new JComboBox<>(new String[]{"High", "Medium", "Low"});
    JLabel lblCategory = new JLabel("Category:");
    cmbCategory = new JComboBox<>(new String[]{"Home", "Work", "Personal"});
    JLabel lblDate = new JLabel("Date (dd-MM-yyyy):");
    txtDate = new JTextField();
    JLabel lblTime = new JLabel("Time (HH:mm:ss):");
    txtTime = new JTextField();
    JButton btnSave = new JButton("Save");
    btnSave.addActionListener(e -> saveTask());
    add(lblName);
    add(txtName);
    add(lblDescription);
    add(txtDescription);
    add(lblPriority);
    add(cmbPriority);
    add(lblCategory);
    add(cmbCategory);
    add(lblDate);
    add(txtDate);
    add(lblTime);
    add(txtTime);
    add(btnSave);
    setDefaultCloseOperation(DISPOSE_ON_CLOSE);
}

// Save task
private void saveTask() {
    String name = txtName.getText().trim();
    String description = txtDescription.getText().trim();
    String priority = (String) cmbPriority.getSelectedItem();
    String category = (String) cmbCategory.getSelectedItem();
    String date = txtDate.getText().trim();
    String time = txtTime.getText().trim();
    if (!name.isEmpty() && !description.isEmpty() && !date.isEmpty() && !time.isEmpty()) {
        String dateTime = date + " " + time;
        parent.addTaskFromDialog(name, description, priority, category, dateTime);
        dispose();
    } else {
        JOptionPane.showMessageDialog(null, "Name, description, date, and time cannot be empty.", "Error", JOptionPane.ERROR_MESSAGE);
    }
}

// Add new category
    private void addNewCategory() {
        String newCategory = txtNewCategory.getText().trim();
        if (!newCategory.isEmpty()) {
            cmbCategory.addItem(newCategory);
            cmbCategory.setSelectedItem(newCategory);
        }
    }
}
// Task edit dialog
private class TaskEditDialog extends JDialog {
    private JTextField txtName;
    private JTextField txtDescription;
    private JComboBox<String> cmbPriority;
    private JComboBox<String> cmbCategory;
    private TextFieldTest parent;
// Constructor
    public TaskEditDialog(JFrame parent, String name, String description, String priority, String category, TextFieldTest parentInstance) {
        super(parent, "Edit Task", true);
        setSize(300, 200);
        setLocationRelativeTo(parent);
        setLayout(new GridLayout(5, 2));
        this.parent = parentInstance;
        JLabel lblName = new JLabel("Name:");
        txtName = new JTextField(name);
        JLabel lblDescription = new JLabel("Description:");
        txtDescription = new JTextField(description);
        JLabel lblPriority = new JLabel("Priority:");
        cmbPriority = new JComboBox<>(new String[]{"High", "Medium", "Low"});
        cmbPriority.setSelectedItem(priority);
        JLabel lblCategory = new

        JLabel("Category:");
        cmbCategory = new JComboBox<>(new String[]{"Home", "Work", "Personal"});
        cmbCategory.setSelectedItem(category);
        JButton btnSave = new JButton("Save");
        btnSave.addActionListener(e -> saveChanges());
        add(lblName);
        add(txtName);
        add(lblDescription);
        add(txtDescription);
        add(lblPriority);
        add(cmbPriority);
        add(lblCategory);
        add(cmbCategory);
        add(btnSave);
        setDefaultCloseOperation(DISPOSE_ON_CLOSE);
        }
        private void saveChanges() {
            String name = txtName.getText().trim();
            String description = txtDescription.getText().trim();
            String priority = (String) cmbPriority.getSelectedItem();
            String category = (String) cmbCategory.getSelectedItem();
            if (!name.isEmpty() && !description.isEmpty()) {
                updateTask(name, description, priority, category);
                dispose();
            } else {
                JOptionPane.showMessageDialog(null, "Name and description cannot be empty.", "Error", JOptionPane.ERROR_MESSAGE);
            }
        }
// Update task
        private void updateTask(String name, String description, String priority, String category) {
            SimpleDateFormat dateFormat = new SimpleDateFormat("dd-MM-yyyy HH:mm:ss");
            String currentDateTime = dateFormat.format(new Date());
    
            JsonObject taskObject = createTaskJson(name, description, priority, category, currentDateTime);
            for (int i = 0; i < tasks.size(); i++) {
                JsonObject existingTask = tasks.get(i);
                if (existingTask.get("description").getAsString().equals(description)) {
                    tasks.set(i, taskObject);
                    break;
                }
            }
            saveTasksToFile();
            refreshTaskDisplay("All", "All");
        }
    }
    // Main method
    public static void main(String[] args) {
        SwingUtilities.invokeLater(() -> TextFieldTest.getInstance());
    }
} 
